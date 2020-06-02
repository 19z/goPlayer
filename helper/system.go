package helper

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type _systemInfo struct {
	FFmpegPath    string
	SystemArch    string
	SystemGoos    string
	TempDir       string
	PathSeparator string
}

var (
	SystemInfo = _systemInfo{}
)

func init() {
	SystemInfo = _systemInfo{
		FFmpegPath:    findFFmpeg(),
		SystemArch:    runtime.GOARCH, //amd64
		SystemGoos:    runtime.GOOS,
		TempDir:       tempDir(),
		PathSeparator: string(os.PathSeparator),
	}
	fmt.Printf("%+v\n", SystemInfo)
}

func findFFmpeg() string {
	//paths := []string{"ffmpeg", "." + string(os.PathSeparator) + "ffmpeg"}
	paths := []string{tempDir() + "/ffmpeg"}
	for _, _path := range paths {
		_path = CommandProgramName(_path)
		if output, err := RunCommand(_path, "-hide_banner"); strings.Index(output, "Video") > 0 {
			println(output)
			return _path
		} else if err != nil {
			println(err.Error())
		}
	}
	path, err := downloadFFmpeg()
	if err != nil {
		log.Fatal(err.Error())
	}
	return path
}

func downloadFFmpeg() (string, error) {
	defaultFileDir := tempDir()
	defaultFilePath := defaultFileDir + "/ffmepeg"
	_ = MakeDir(defaultFileDir)
	e := errors.New("please install ffmpeg")
	result := struct {
		Release []string `json:"release"`
		Lgpl    []string `json:"lgpl"`
	}{}
	err := httplib.Get("https://ffmpeg.zeranoe.com/builds/builds.json").SetTimeout(15*time.Second, 10*time.Second).ToJSON(&result)
	if err != nil {
		return "", e
	}
	println("try download ffmpeg " + result.Lgpl[0] + "...")
	if runtime.GOOS == "windows" {
		filename := "ffmpeg-" + result.Lgpl[0] + "-win64-static.zip"
		url := "https://ffmpeg.zeranoe.com/builds/win64/static/" + filename
		zipFile := PathJoin(defaultFileDir, filename)
		//defer os.Remove(zipFile)
		if !FileExist(zipFile) {
			if DownloadFile(zipFile, url) != nil {
				return "", e
			}
		}
		// 解压出 ffmpeg
		zipReader, err := zip.OpenReader(zipFile)
		if err != nil {
			println(err.Error())
			return "", e
		}
		defer zipReader.Close()
		for _, f := range zipReader.File {
			fname := Basename(f.Name)
			if fname == "ffmpeg" || fname == "ffmpeg.exe" {
				if inFile, err := f.Open(); err == nil {
					defer inFile.Close()
					if outFile, err := os.OpenFile(defaultFileDir+"/"+fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777); err == nil {
						defer outFile.Close()
						_, err = io.Copy(outFile, inFile)
						break
					}
				}
			}
		}
	}
	return defaultFilePath, nil
}

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CommandProgramName(name string) string {
	if os.PathSeparator == '\\' {
		return strings.ReplaceAll(name, "/", "\\")
	}
	return name
}

func RunCommand(name string, args ...string) (string, error) {
	var outInfo bytes.Buffer
	name = CommandProgramName(name)
	println(name, args[0])
	cmd := exec.Command(name, args...)
	cmd.Stderr = &outInfo
	err := cmd.Run()
	return outInfo.String(), err
}

func tempDir() string {
	return "./.temp"
}

func MakeDir(filePath string) error {
	if !FileExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func FileExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

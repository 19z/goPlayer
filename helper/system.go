package helper

import "os"

func FindFFmpeg() string {
	return "C:\\MineProgram\\ffmpeg\\bin\\ffmpeg.exe"
}
func TempDir() string {
	return "./.temp/"
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

package drivers

import (
	"github.com/astaxie/beego/context"
	"goPlayer/helper"
	"io/ioutil"
	"mime"
	"os"
	"path"
	"strconv"
	"strings"
)

type LocalDriver struct {
	Id   int
	Path string
}

func (d LocalDriver) DirList(p string) []FileItem {
	p = strings.Trim(p, "/\\")
	//println(d.Path, p)
	files, err := ioutil.ReadDir(path.Join(d.Path, p))
	//println(path.Join(d.Path, p))
	if err != nil {
		return nil
	}
	result := []FileItem{}
	for _, file := range files {
		item := FileItem{
			IsDirectory: file.IsDir(),
			Directory:   p,
			Name:        file.Name(),
			Size:        file.Size(),
			MimeType:    "",
			ModTime:     file.ModTime(),
			driver:      d,
		}
		if !item.IsDirectory {
			item.MimeType = mime.TypeByExtension(p)
		}
		result = append(result, item)
	}
	return result
}

func (d LocalDriver) GetPreviewUrl(item FileItem) string {
	if !item.IsDirectory {
		if helper.IsVideo(item.Name) {
			return "/video/" + strconv.Itoa(d.Id) + "/" + item.GetFullPath()
		}
	}
	return "/files/" + strconv.Itoa(d.Id) + "/" + item.GetFullPath()
}

func (d LocalDriver) GetDownloadUrl(item FileItem) string {
	return "/files/" + strconv.Itoa(d.Id) + "/" + item.GetFullPath()
}

func (d LocalDriver) GetFileItem(p string) (FileItem, error) {
	p = strings.Trim(p, "/\\")
	fullPath := strings.TrimRight(d.Path, "/\\") + "/" + p
	stat, err := os.Stat(fullPath)
	if err != nil {
		return FileItem{}, err
	}
	//todo : return FileItem
	return FileItem{
		IsDirectory: stat.IsDir(),
		Directory:   helper.Dirname(p),
		Name:        stat.Name(),
		Size:        stat.Size(),
		MimeType:    mime.TypeByExtension("." + helper.ArrayEnd(strings.Split(stat.Name(), "."))),
		driver:      d,
	}, nil
}

func (d LocalDriver) SendFile(file *FileItem, c *context.Context) {
	fullPath := path.Join(d.Path, file.Directory, file.Name)
	//println(fullPath)
	c.Output.Download(fullPath, file.Name)
}

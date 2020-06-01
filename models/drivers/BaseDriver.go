package drivers

import (
	"github.com/astaxie/beego/context"
	"path"
	"time"
)

type FileItem struct {
	IsDirectory bool
	Directory   string
	Name        string
	Size        int64
	ModTime       time.Time
	MimeType    string
	Children    *[]FileItem
	driver      BaseDriver
}

type BaseDriver interface {
	DirList(dirname string) []FileItem         // 获取所有子文件
	GetPreviewUrl(item FileItem) string        // 获取预览地址
	GetDownloadUrl(item FileItem) string       //获取下载地址
	GetFileItem(item string) (FileItem, error) //获取下载地址
	SendFile(*FileItem, *context.Context)      // 下载文件
}

func (c FileItem) GetFullPath() string {
	return path.Join(c.Directory, c.Name)
}

func (c FileItem) GetPreviewUrl() string {
	if c.driver != nil {
		return c.driver.GetPreviewUrl(c)
	}
	return ""
}

func (c FileItem) GetDownloadUrl() string {
	if c.driver != nil {
		return c.driver.GetDownloadUrl(c)
	}
	return ""
}

func (f FileItem) SendFile(c *context.Context) {
	if f.driver != nil {
		f.driver.SendFile(&f, c)
	}
}

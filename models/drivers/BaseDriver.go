package drivers

import "strings"

type FileItem struct {
	IsDirectory bool
	Directory   string
	Name        string
	Size        int64
	MimeType    string
	Children    *[]FileItem
	driver      BaseDriver
}

func (c FileItem) GetFullPath() string {
	dir := strings.Trim(c.Directory, "/\\")
	if len(dir) > 0 {
		return dir + "/" + c.Name
	} else {
		return c.Name
	}
}

func (c FileItem) GetPreviewUrl() string {
	if c.driver != nil {
		return c.driver.GetDownloadUrl(c)
	}
	return ""
}

func (c FileItem) GetDownloadUrl() string {
	if c.driver != nil {
		return c.driver.GetDownloadUrl(c)
	}
	return ""
}

type BaseDriver interface {
	DirList(dirname string) []FileItem   // 获取所有子文件
	GetPreviewUrl(item FileItem) string  // 获取预览地址
	GetDownloadUrl(item FileItem) string //获取下载地址
	GetFileItem(item string) FileItem    //获取下载地址
}

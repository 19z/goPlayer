package drivers

import (
	"io/ioutil"
	"mime"
	"os"
	"strconv"
	"strings"
)

type LocalDriver struct {
	Id   int
	Path string
}

func (d LocalDriver) DirList(path string) []FileItem {
	path = strings.Trim(path, "/\\")
	files, err := ioutil.ReadDir(strings.TrimRight(d.Path, "/\\") + "/" + path)
	if err != nil {
		return nil
	}
	result := []FileItem{}
	for _, file := range files {
		item := FileItem{
			IsDirectory: file.IsDir(),
			Directory:   path,
			Name:        file.Name(),
			Size:        file.Size(),
			MimeType:    "",
			driver:      d,
		}
		if !item.IsDirectory {
			item.MimeType = mime.TypeByExtension(path)
		}
		result = append(result, item)
	}
	return result
}

func (d LocalDriver) GetPreviewUrl(item FileItem) string {
	if !item.IsDirectory {
		if len(item.MimeType) > 5 && item.MimeType[0:5] == "video" {
			return "/video/" + strconv.Itoa(d.Id) + "/" + item.GetFullPath()
		} else {
			return ""
		}
	}
	return "/preview/" + strconv.Itoa(d.Id) + "/" + item.GetFullPath()
}

func (d LocalDriver) GetDownloadUrl(item FileItem) string {
	return "/files/" + strconv.Itoa(d.Id) + "/" + item.GetFullPath()
}

func (d LocalDriver) GetFileItem(path string) *FileItem {
	path = strings.Trim(path, "/\\")
	fullPath := strings.TrimRight(d.Path, "/\\") + "/" + path
	_, err := os.Stat(fullPath)
	if err != nil {
		return nil
	}
	//todo : return FileItem
	return nil
}

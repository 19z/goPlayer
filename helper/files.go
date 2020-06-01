package helper

import (
	"fmt"
	"mime"
	"path"
	"strconv"
	"strings"
	"time"
)

func Dirname(filepath string) string {
	return path.Dir(strings.TrimRight(strings.ReplaceAll(filepath, "\\", "/"), "/"))
}

func PathJoin(elem ...string) string {
	return path.Join(elem...)
}

func IsVideo(p string) bool {
	ext := ArrayEnd(strings.Split(p, "."))
	videos := []string{"m3u8", "rmvb", "mp4"}
	if InArray(videos, ext) {
		return true
	}
	m := mime.TypeByExtension(ext)
	if len(m) > 5 && m[0:5] == "video" {
		return true
	}
	return false
}

func SplatUrl(url string) (*int, string, error) {
	splatList := strings.Split(url, "/")
	driverId, err := strconv.Atoi(splatList[0])
	splat := strings.Join(ArrayFilter(splatList[1:]), "/")

	if err != nil {
		return nil, splat, err
	}
	return &driverId, splat, nil
}

func FileSizeFormat(size int64) string {
	f := "BKMGT"
	v := float64(size)
	i := 0;
	for v > 1024 {
		i++
		v = v / 1024.0
	}
	return fmt.Sprintf("%.2f", v) + " " + f[i:i+1];
}

func TimeFormat(t time.Time) string {
	now := time.Now()
	i := now.Unix() - t.Unix()
	before := i < 0;
	if before {
		i = -i
	}
	if i < 60*1000 && !before {
		return "刚刚";
	} else if i < 60*1000 {
		return ""
	}
	return t.Format("2006/02/01 15:04:05")
}

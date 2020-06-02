package ffmpeg

import (
	"errors"
	"goPlayer/helper"
)

type FFmpeg struct {
	Input  string
	Output string
}

func FFmpegsNew(file string, outputDir string) (FFmpeg, error) {
	if helper.SystemInfo.FFmpegPath == "" {
		return FFmpeg{}, errors.New("not ffmpeg")
	}
	if helper.FileExist(file) {
		return FFmpeg{Input: file, Output: outputDir}, nil
	} else {
		return FFmpeg{}, errors.New("file not exists")
	}
}
func (f FFmpeg) GetThumbnail() (string, error) {
	savePath := helper.PathJoin(f.Output, "thumbnail.jpg")
	outPut, err := helper.RunCommand(helper.SystemInfo.FFmpegPath, "-hide_banner", "-i", f.Input, "-ss", "100", "-y", "-t", "0.001", "-f", "image2", savePath)
	if outPut != "" {
		return savePath, nil
	} else {
		return savePath, err
	}
}

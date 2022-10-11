package utils

import (
	"os"
	"strconv"
	"time"

	"golang.design/x/clipboard"
)

func SavePngFromClipboard() (bool, string) {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	b2 := clipboard.Read(clipboard.FmtImage)
	i := len(b2)
	if i != 0 {
		i2 := time.Now().Local().UnixMilli()
		unix := strconv.FormatInt(i2, 10)
		s, _ := os.UserHomeDir()
		imgPath := s + "/GTools/" + unix + ".png"
		f, err := os.Create(imgPath)
		f.Write(b2)
		defer f.Close()
		if err != nil {
			return false, err.Error()
		}
		return true, imgPath
	} else {
		return false, "系统粘贴板中没有图片！"
	}
}

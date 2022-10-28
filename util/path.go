package util

import (
	"os"
)

func CheckFileExist(path string) (bool, string) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err.Error()
	}
	if fi.IsDir() {
		return false, "路径为文件夹"
	} else {
		return true, ""
	}
}

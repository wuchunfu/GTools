package utils

import (
	"os"
	"strings"
)

// 获取目录下所有的MarkDown文件（不递归）
func GetMdFileList(name string) []map[string]string {
	filePathList := make([]map[string]string, 0)
	de, err := os.ReadDir(name)
	if err != nil {
		return filePathList
	}
	for _, v := range de {
		s := strings.Split(v.Name(), ".")
		if s[len(s)-1] == "md" {
			mapPath := make(map[string]string)
			fname := v.Name()
			mapPath["fname"] = fname
			mapPath["fpath"] = name + "/" + v.Name()
			filePathList = append(filePathList, mapPath)
		}
	}
	return filePathList
}

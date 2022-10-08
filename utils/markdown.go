package utils

import (
	"fmt"
	"io/fs"
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

func GetFileContent(path string) string {
	fmt.Printf("path: %v\n", path)
	var resp Resp[string]
	fi, err := os.Stat(path)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		return resp.toJson()
	}
	if fi.IsDir() {
		resp.Code = 300
		resp.Msg = "不能读取文件夹"
		return resp.toJson()
	}
	b, err := os.ReadFile(path)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		return resp.toJson()
	} else {
		resp.Code = 200
		resp.Data = string(b)
		return resp.toJson()
	}
}

func SaveFileContent(path string, content string) string {
	var resp Resp[string]
	fi, err := os.Stat(path)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		return resp.toJson()
	}
	if fi.IsDir() {
		resp.Code = 500
		resp.Msg = "内容写入失败"
		return resp.toJson()
	}
	err2 := os.WriteFile(path, []byte(content), fs.FileMode(os.O_RDWR|os.O_TRUNC))
	if err2 != nil {
		resp.Code = 500
		resp.Msg = err2.Error()
		return resp.toJson()
	} else {
		resp.Code = 200
		return resp.toJson()
	}
}

package utils

import (
	"encoding/json"
	"os"
	"strings"
)

// 获取目录下所有的MarkDown文件（不递归）
func GetMdFileList(name string) string {
	ret := make(map[string]string)
	filePathList := make([]map[string]string, 0)
	de, err := os.ReadDir(name)
	if err != nil {
		ret["code"] = "0"
		ret["msg"] = err.Error()
		ret["data"] = ""
		json2str, _ := json.Marshal(ret)
		s := string(json2str)
		return s
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
	mdArrBit, _ := json.Marshal(filePathList)
	mdArrStr := string(mdArrBit)
	ret["code"] = "1"
	ret["msg"] = "OK"
	ret["data"] = mdArrStr
	resultBit, _ := json.Marshal(ret)
	resultStr := string(resultBit)
	return resultStr
}

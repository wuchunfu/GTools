package utils

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

type Resp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func (a *Resp[T]) toJson() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// 将文件夹路径添加到json数据库中
func AddDirPathToJdb(dirPath string) string {
	var resp Resp[string]

	// 判断路径是否为文件夹
	fi, err := os.Stat(dirPath)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		return resp.toJson()
	}
	if !fi.IsDir() {
		resp.Code = 500
		resp.Msg = "非文件夹路径"
		return resp.toJson()
	}

	// 初始化数据路径
	b, path, err := InitJdb("path.json")
	if !b {
		// 路径数据库初始化失败
		resp.Code = 500
		resp.Msg = err.Error()
		return resp.toJson()
	}
	b2, err2 := os.ReadFile(path)
	if err2 != nil {
		resp.Code = 500
		resp.Msg = "err2" + err2.Error()
		return resp.toJson()
	}
	pathMap := make(map[string]string)
	if len(b2) == 0 {
		pathMap[dirPath] = dirPath
		fmt.Printf("pathMap: %v\n", pathMap)
		b3, _ := json.MarshalIndent(pathMap,"","  ")
		err4 := os.WriteFile(path, b3, fs.FileMode(os.O_RDWR|os.O_CREATE))
		if err4 != nil {
			resp.Code = 500
			resp.Msg = "err4" + err4.Error()
			return resp.toJson()
		}
		resp.Code = 200
		resp.Data = string(b3)
		return resp.toJson()
	} else {
		err3 := json.Unmarshal(b2, &pathMap)
		if err3 != nil {
			resp.Code = 500
			resp.Msg = "err3" + err3.Error()
			return resp.toJson()
		}
		fmt.Printf("pathMap[dirPath]: %v\n", pathMap[dirPath])
		pathMap[dirPath] = dirPath
		b3, _ := json.MarshalIndent(pathMap,"","  ")
		err4 := os.WriteFile(path, b3, fs.FileMode(os.O_RDWR|os.O_TRUNC))
		if err4 != nil {
			resp.Code = 500
			resp.Msg = "err4" + err4.Error()
			return resp.toJson()
		}
		resp.Code = 200
		resp.Data = string(b3)
		return resp.toJson()
	}
}

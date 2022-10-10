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
	pathData, err2 := os.ReadFile(path)
	if err2 != nil {
		resp.Code = 500
		resp.Msg = "path数据异常: " + err2.Error()
		return resp.toJson()
	}
	pathMap := make(map[string][]map[string]string) // Map 中key为目录，value为文件路径列表
	// 获取目录下的md文件路径列表
	if len(pathData) != 0 { // 如果数据库中没有数据则直接使用填充数据，否则取出后再进行修改
		json.Unmarshal(pathData, &pathMap)
	}
	mdList := GetMdFileList(dirPath)
	pathMap[dirPath] = mdList
	b3, _ := json.MarshalIndent(pathMap, "", "  ")
	err4 := os.WriteFile(path, b3, fs.FileMode(os.O_RDWR|os.O_CREATE))
	if err4 != nil {
		resp.Code = 500
		resp.Msg = "err4" + err4.Error()
		return resp.toJson()
	}
	resp.Code = 200
	resp.Data = string(b3)
	return resp.toJson()
}

// 获取路径所有存储的文件夹路径数据及Markdown文档列表
func GetPathData() (string) {
	// 从指定的数据json中取出所有数据
	b, b2 := GetJdbContent("/path.json")
	var resp Resp[string]
	if b {
		resp.Code = 200
		resp.Msg = "OK"
		resp.Data = string(b2)
		return resp.toJson()
	} else {
		resp.Code = 500
		resp.Msg = string(b2)
		resp.Data = ""
		return resp.toJson()
	}
}

func DelDirPathFromJdb(path string) string  {
	fmt.Printf("path: %v\n", path)
	var resp Resp[string]
	resp.Code = 500
	fi, err := os.Stat(path)
	if err != nil {
		resp.Msg = err.Error()
		return resp.toJson()
	}
	if !fi.IsDir() {
		resp.Msg = "非文件夹路径"
		return resp.toJson()
	}
	// 从指定的数据json中取出所有数据
	b, data := GetJdbContent("/path.json")
	if !b {
		resp.Msg = "数据文件读取失败"
		return resp.toJson()
	}
	var pathMap = make(map[string][]map[string]string)
	err2 := json.Unmarshal(data, &pathMap)
	if err2 != nil {
		resp.Msg = err2.Error()
		return resp.toJson()
	}
	delete(pathMap, path)
	newData, _ := json.MarshalIndent(pathMap, "", "  ")
	// 初始化数据路径
	_, jpath, _ := InitJdb("path.json")
	err3 := os.WriteFile(jpath, newData, fs.FileMode(os.O_RDWR|os.O_TRUNC))
	if err3 != nil {
		resp.Msg = err3.Error()
		return resp.toJson()
	}
	resp.Code = 200
	return resp.toJson()
}

func CheckFileExist(path string) (bool, string)  {
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
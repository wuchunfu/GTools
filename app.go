package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// 获取文件夹下的所有markdown文件
func (a *App) GetFileList(name string) string {
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

// 获取md文档的全部内容
func (a *App) GetMdContent(name string) string {
	ret := make(map[string]string)
	b, err := os.ReadFile(name)
	if err != nil {
		ret["code"] = "0"
		ret["msg"] = err.Error()
		ret["data"] = ""
		json2str, _ := json.Marshal(ret)
		s := string(json2str)
		return s
	}
	content := string(b)
	ret["code"] = "1"
	ret["msg"] = "OK"
	ret["data"] = content
	json2str, _ := json.Marshal(ret)
	s := string(json2str)
	return s
}

func (a *App) UploadImgByPicgo(name string) string {
	b, _ := exec.Command("picgo", "upload").Output()
	resultStr := string(b)
	s := strings.Trim(resultStr, " ")
	isContain := strings.ContainsAny(s, "https://")
	if isContain {
		s2 := strings.Split(s, "[PicGo SUCCESS]:")
		fmt.Printf("s: %v\n", s2[1])
		return s2[1]
	} else {
		fmt.Printf("error")
	}
	return fmt.Sprintf("Hello %s, I am Xiao!", name)
}

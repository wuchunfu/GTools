package main

import (
	"changeme/utils"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	fmt.Printf("name: %v\n", name)
	return fmt.Sprintf("Hello %s, It's show time!", name)
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
	os.Chdir("/Users/pixiao/GTools")
	b, err := exec.Command("sh", "-c", "./picgo.sh").Output()
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
	// resultStr := string(b)
	// fmt.Printf("resultStr: %v\n", resultStr)
	// s := strings.Trim(resultStr, " ")
	// isContain := strings.ContainsAny(s, "https://")
	// if isContain {
	// 	s2 := strings.Split(s, "[PicGo SUCCESS]:")
	// 	fmt.Printf("s: %v\n", s2[1])
	// 	return s2[1]
	// } else {
	// 	return resultStr
	// }
}

func (a *App) OpenUrl(url string) {
	fmt.Printf("url: %v\n", url)
	runtime.BrowserOpenURL(a.ctx, "http://www.baidu.com")
}

func (a *App) GetWd() string {
	os.Chdir("/Users/pixiao/GTools")
	dir, _ := os.Getwd()
	return dir
}

func (a *App) AddDirPath(path string) string {
	return utils.AddDirPathToJdb(path)
}

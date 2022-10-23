package gtools

import (
	"changeme/util"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) OpenSelectFileWindow() *util.Resp {
	options := runtime.OpenDialogOptions{}
	options.Title = "请选择文件"
	options.CanCreateDirectories = true
	s, err := runtime.OpenFileDialog(a.ctx, options)
	if err != nil {
		panic(err.Error())
	}
	data := make(map[string]string)
	data["res"] = s
	util.Success(data)
	fmt.Printf("Open: %v\n", s)
	runtime.LogPrint(a.ctx, s)
	return util.Success(data)
}

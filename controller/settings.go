package gtools

import (
	"gtools/configs"
	"gtools/util"
	"fmt"
	"os"
	"os/user"
)

func (a *App) CleanWebKitCache() *util.Resp {
	// 直接将MacOS上对应的Cache文件夹删除
	userInfo, err := user.Current()
	if err != nil {
		return util.Error(err.Error())
	}
	cachePath := fmt.Sprintf(configs.WebkitPath, userInfo.HomeDir)
	// 判断文件夹是否存在
	if _, err = os.Stat(cachePath); err != nil {
		return util.Error(err.Error())
	}
	// 删除缓存文件夹及内部文件
	if err = os.RemoveAll(cachePath); err != nil {
		return util.Error(err.Error())
	}
	a.Log.Info(fmt.Sprintf("缓存已清理，缓存路径为：%s", cachePath))
	return util.Ok()
}

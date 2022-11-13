package gtools

import (
	"fmt"
	"gtools/configs"
	"gtools/internal"
	"gtools/util"
	"os"
	"os/user"
)

// 页面加载是获取系统配置
func (a *App) GetConfigOnounted() *util.Resp {
	return util.Success(a.ConfigMap)
}

func (a *App) UpdateConfigItem(item internal.ConfigItem) *util.Resp {
	var value string = item.Value
	var search = internal.ConfigItem {
		Name: item.Name,
		Type: item.Type,
	}
	if _, err := a.Db.Where("name = ?", item.Name).Where("type = ?", item.Type).Get(&search); err != nil {
		a.Log.Error(fmt.Sprintf(configs.GetConfigItemErr, item.Name, err.Error()))
		return util.Error(err.Error())
	}
	// 值没有变化
	if value == search.Value {
		return util.Ok()
	}
	search.Value = value
	// 值有变化，则进行更新，同时需要配置上下文信息
	if _, err := a.Db.ID(search.Id).AllCols().Update(&search); err != nil {
		a.Log.Error(fmt.Sprintf(configs.UpdateConfigItemErr, item.Name, err.Error()))
		return util.Error(err.Error())
	}
	a.ConfigMap = a.GetConfigMap()
	return util.Success(a.ConfigMap)
}

// 获取系统配置
func (a *App) GetConfigMap() map[string]map[string]string {
	configMap := make(map[string]map[string]string)
	for _, ctype := range configs.ConfigTypes {
		list := make([]internal.ConfigItem, 0)
		if err := a.Db.Where("type = ?", ctype).Find(&list); err != nil {
			a.Log.Error(fmt.Sprintf(configs.GetConfigListErr, err.Error()))
			panic(err.Error())
		}
		if len(list) == 0 {
			a.initConfigData(ctype)
		} else {
			typeMap := make(map[string]string, 0)
			for _, v := range list {
				typeMap[v.Name] = v.Value
			}
			configMap[ctype] = typeMap
		}
	}
	return configMap
}

// 清理Mac上的Webkit缓存
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

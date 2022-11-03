package gtools

import (
	"changeme/configs"
	"changeme/internal"
	"changeme/util"
)

func (a *App) AddCmdItem(item internal.CmdItem) *util.Resp {
	if _, err := a.Db.Insert(&item); err != nil {
		a.Log.Error(configs.AddCmdItemErr, item.Name, err.Error())
		return util.Error(err.Error())
	}
	return a.GetCmdItemList()
}

func (a *App) GetCmdItemList() *util.Resp {
	list := make([]internal.CmdItem, 0)
	if err := a.Db.Find(&list); err != nil {
		return util.Error(err.Error())
	}
	return util.Success(list)
}

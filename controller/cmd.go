package gtools

import (
	"gtools/util"
	"fmt"
	"gtools/configs"
	"gtools/internal"
	"os/exec"
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

func (a *App) DelCmdItem(item internal.CmdItem) *util.Resp {
	if _, err := a.Db.Delete(&item); err != nil {
		a.Log.Error(fmt.Sprintf(configs.DelCmdItemErr, item.Name, err.Error()))
		return util.Error(err.Error())
	}
	return util.Ok()
}

func (a *App) CmdHandler(item internal.CmdItem) *util.Resp {
	var cmdStr string
	// 更新状态
	if item.State == 0 {
		cmdStr = item.Start
		item.State = 1
	} else {
		cmdStr = item.Stop
		item.State = 0
	}
	fmt.Printf("item: %v\n", item)
	if _, err := exec.Command(item.Type, cmdStr).Output(); err != nil {
		a.Log.Error(fmt.Sprintf(configs.CmdItemHandlerErr, item.Name, err.Error()))
		return util.Error(err.Error())
	}
	if _, err := a.Db.AllCols().Update(&item); err != nil {
		a.Log.Error(fmt.Sprintf(configs.UpdateCmdItemErr, item.Name, err.Error()))
		return util.Error(err.Error())
	}
	return a.GetCmdItemList()
}

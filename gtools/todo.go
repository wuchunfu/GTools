package gtools

import (
	"changeme/configs"
	"changeme/internal"
	"changeme/util"
	"fmt"
)

func (a *App) AddTodoItem(item internal.TodoItem) *util.Resp {
	_, err := a.Db.InsertOne(&item)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.AddTodoItemErr, item.Title, err.Error()))
		return util.Error(err.Error())
	}
	return a.GetTodoList()
}

func (a *App) GetTodoList() *util.Resp {
	itemList := make([]internal.TodoItem, 0)
	err := a.Db.Desc("date").Find(&itemList)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.GetTodoListErr, err.Error()))
		return util.Error(err.Error())
	}
	return util.Success(itemList)
}

func (a *App) DelTodoItem(item internal.TodoItem) *util.Resp {
	_, err := a.Db.ID(item.Id).Delete(&item)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.DelTodoItemErr, item.Title, err.Error()))
		return util.Error(err.Error())
	}
	return a.GetTodoList()
}

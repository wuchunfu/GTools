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
	doneItemList := make([]internal.TodoItem, 0)
	todoItemList := make([]internal.TodoItem, 0)
	err := a.Db.Desc("date").Where("done = ?", 1).Find(&doneItemList)
	err2 := a.Db.Desc("date").Where("done = ?", 0).Find(&todoItemList)
	if err != nil{
		a.Log.Error(fmt.Sprintf(configs.GetTodoListErr, err.Error()))
		return util.Error(err.Error())
	}
	if err2 != nil{
		a.Log.Error(fmt.Sprintf(configs.GetTodoListErr, err2.Error()))
		return util.Error(err2.Error())
	}
	resultMap := make(map[string][]internal.TodoItem, 0)
	resultMap["done"] = doneItemList
	resultMap["todo"] = todoItemList
	return util.Success(resultMap)
}

func (a *App) DelTodoItem(item internal.TodoItem) *util.Resp {
	_, err := a.Db.Delete(&item)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.DelTodoItemErr, item.Title, err.Error()))
		return util.Error(err.Error())
	}
	return a.GetTodoList()
}

func (a *App) GetTodoItem(id int) *util.Resp {

	return util.Ok()
}

func (a *App) UpdateTodoItem(item internal.TodoItem) *util.Resp {
	_, err := a.Db.ID(item.Id).AllCols().Update(&item)
	if err != nil {
		return util.Error(err.Error())
	}
	return a.GetTodoList()
}

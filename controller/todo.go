package gtools

import (
	"fmt"
	"gtools/configs"
	"gtools/internal"
	"gtools/util"
	"math"
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
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.GetTodoListErr, err.Error()))
		return util.Error(err.Error())
	}
	if err2 != nil {
		a.Log.Error(fmt.Sprintf(configs.GetTodoListErr, err2.Error()))
		return util.Error(err2.Error())
	}
	doneItemNum := len(doneItemList)
	todoItemNUm := len(todoItemList)
	totaltotalItemNum := doneItemNum + todoItemNUm
	var rate float64 = 0
	if totaltotalItemNum != 0 {
		frate := float64(doneItemNum) / float64(totaltotalItemNum)
		rate = math.Floor(frate * 100)
	}

	resultMap := make(map[string]interface{}, 0)
	resultMap["doneList"] = doneItemList
	resultMap["todoList"] = todoItemList
	resultMap["rate"] = rate
	resultMap["todonum"] = todoItemNUm // 未完成待办事项
	resultMap["donenum"] = doneItemNum
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

func (a *App) DelTodoItemById(item internal.TodoItem) *util.Resp {
	_, err := a.Db.ID(item.Id).Delete(&item)
	if err != nil {
		return util.Error(err.Error())
	}
	return a.GetTodoList()
}

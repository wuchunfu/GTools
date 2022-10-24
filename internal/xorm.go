package internal

import (
	_ "github.com/mattn/go-sqlite3"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type TodoItem struct {
	ID      int    `json:"id" xorm:"autoincr"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewXormEngine(dbPath string) *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", dbPath)
	if err != nil {
		panic("数据库初始化失败: " + err.Error())
	}
	// 设置连接池大小
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10)
	engine.SetMapper(names.GonicMapper{}) // 名称映射规则 驼峰式
	engine.Sync2(new(TodoItem))
	return engine
}

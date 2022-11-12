package internal

import (
	"time"

	_ "github.com/mattn/go-sqlite3"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func NewXormEngine(dbPath string) *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", dbPath)
	// 引入 _ "github.com/go-sql-driver/mysql" 即可使用mysql进行数据存储， 配置如下
	// engine, err := xorm.NewEngine("mysql", "root:root@/gtools?charset=utf8")
	if err != nil {
		panic("数据库初始化失败: " + err.Error())
	}
	// 设置时区和数据库时区
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.DatabaseTZ, _ = time.LoadLocation("Asia/Shanghai")

	// 控制台输出SQL语句
	engine.ShowSQL(true)

	// 设置连接池大小
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10)
	engine.SetMapper(names.GonicMapper{}) // 名称映射规则 驼峰式
	engine.Sync2(new(TodoItem), new(MdPath), new(CmdItem), new(ConfigItem))
	return engine
}

package internal

import (
	"time"

	_ "github.com/mattn/go-sqlite3"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// 待办事项表
type TodoItem struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title" xorm:"varchar(100) default('') notnull"`    // 标题
	Content    string    `json:"content" xorm:"varchar(1024) default('') notnull"` // 内容
	Tags       string    `json:"tags" xorm:"varchar(200) default('') notnull"`     // 标签
	Date       time.Time `json:"date" xorm:"created"`                              // 日期
	HasContent bool      `json:"hasContent" xorm:"default(0) notnull"`             // 是否有内容
	Done       bool      `json:"done" xorm:"default(0) notnull"`                   // 是否已完成
	Level      int8      `json:"level" xorm:"default(0) notnull"`                  // 重要等级
	Expired    time.Time `json:"expired" xorm:"created"`                           // 事项到期时间
	Updated    time.Time `json:"updated" xorm:"updated"`                           // 更新时间
}

// markdown文件
type MdPath struct {
	Id      int64     `json:"id"`
	Path    string    `json:"path" xorm:"varchar(100) unique"` // 路径
	Type    int8      `json:"type" xorm:"default(0) notnull"`  // 0-文件夹 1-文件
	Fname   string    `json:"fname" xorm:"varchar(100)"`       // 文件名称
	Created time.Time `json:"created" xorm:"created"`          // 创建时间
}

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
	// 设置连接池大小
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10)
	engine.SetMapper(names.GonicMapper{}) // 名称映射规则 驼峰式
	engine.Sync2(new(TodoItem), new(MdPath))
	return engine
}

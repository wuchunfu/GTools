package internal

import "time"

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

// cmd快捷指令
type CmdItem struct {
	Id      int64     `json:"id"`                                             // ID
	Name    string    `json:"name" xorm:"varchar(100) default('') notnull"`   // 名称
	Type    string    `json:"type" xorm:"varchar(200) default('') notnull"`   // 执行类型
	State   int8      `json:"state" xorm:"default(0) notnull"`                // 状态 0-停止 1-运行
	Port    string    `json:"port" xorm:"varchar(20) default(0) notnull"`     // 端口号
	Start   string    `json:"start" xorm:"varchar(1024) default('') notnull"` // 启动指令
	Stop    string    `json:"stop" xorm:"varchar(1024) default('') notnull"`  // 停止指令
	Created time.Time `json:"created" xorm:"created"`                         // 创建时间
}

// 系统配置项
type ConfigItem struct {
	Id      int64     `json:"id"`                                            // ID
	Name    string    `json:"name" xorm:"varchar(200) default('') notnull"`  // 配置项名称
	Value   string    `json:"value" xorm:"varchar(200) default('') notnull"` // 配置项值
	Type    string    `json:"type" xorm:"varchar(100) default('') notnull"`  // 配置项类型
	State   int8      `json:"state" xorm:"default(0) notnull"`               // 启用状态
	Created time.Time `json:"created" xorm:"created"`                        // 创建时间
}

// 图床图片
type ImageItem struct {
	Id      int64     `json:"id"`                                 // ID
	Path    string    `json:"path" xorm:"default('') notnull"`    // 图片路径
	Type    string    `json:"type" xorm:"default('') notnull"`    // 图片类型-本地/阿里OSS/腾讯OSS/sm.ms
	Artname string    `json:"artname" xorm:"default('') notnull"` // 图片所在文章名称
	Artpath string    `json:"artpath" xorm:"default('') notnull"` // 图片所在文章路径
	Created time.Time `json:"created" xorm:"created"`             // 创建时间
}

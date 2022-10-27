package gtools

import (
	"changeme/configs"
	"changeme/internal"
	"changeme/util"
	"context"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

// App 结构体
type App struct {
	ctx     context.Context
	Log     *logrus.Logger
	Db      *xorm.Engine
	LogFile string
	DBFile  string
	AliOSS  *oss.Client
}

// NewApp 创建一个新的 App 应用程序结构体
func NewApp() *App {
	return &App{}
}

// startup 在应用程序启动时调用。 上下文被保存
// 所以我们可以调用 runtime 方法
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx

	// 获取GTools数据文件夹路径
	confDir := util.GetConfigDir()

	// 初始化logrus
	a.LogFile = fmt.Sprintf(configs.LogFile, confDir)
	a.Log = internal.NewLogger(a.LogFile)

	// 初始化xorm
	a.DBFile = fmt.Sprintf(configs.DBFile, confDir)
	a.Db = internal.NewXormEngine(a.DBFile)

	// 初始化阿里云OSS客户端
	a.AliOSS = internal.NewOssClient()
}

// OnBeforeClose
func (a *App) OnBeforeClose(ctx context.Context) bool {
	// 关闭xorm连接
	a.Db.Close()
	// 返回 true 将阻止程序关闭
	return false
}

func (a *App) StartTomcat() string {
	return util.RunCmd()
}

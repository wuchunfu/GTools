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

	// 配置logrus
	a.LogFile = fmt.Sprintf(configs.LogFile, confDir)
	a.Log = internal.NewLogger(a.LogFile)

	// 配置xorm
	a.DBFile = fmt.Sprintf(configs.DBFile, confDir)
	a.Db = internal.NewXormEngine(a.DBFile)

	// 配置阿里云OSS
	a.AliOSS = internal.NewOssClient()
}

func (a *App) StartTomcat() string {
	return util.RunCmd()
}

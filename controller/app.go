package gtools

import (
	"context"
	"fmt"
	"gtools/configs"
	"gtools/internal"
	"gtools/util"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
	"xorm.io/xorm"
)

// App 结构体
type App struct {
	ctx       context.Context
	Log       *logrus.Logger
	Db        *xorm.Engine
	LogFile   string
	DBFile    string
	ConfigMap map[string]map[string]string
	AliOSS    *oss.Client
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

	// 获取系统配置文件
	a.ConfigMap = a.GetConfigMap()

	// 初始化图床配置
	a.initImgBed()

	// 初始化系统剪贴板工具
	a.initClipboard()
}

// OnBeforeClose
func (a *App) OnBeforeClose(ctx context.Context) bool {
	// 关闭xorm连接
	a.Db.Close()
	// 返回 true 将阻止程序关闭
	return false
}

// OnDOMReady
func (a *App) OnDOMReady(ctx context.Context) {
	// 启动一个监听事件
	runtime.EventsOn(a.ctx, "test", func(optionalData ...interface{}) {
		a.Log.Info(optionalData...)
	})
}

// 初始化图床配置
func (a *App) initImgBed() {
	switch a.ConfigMap["imgbed"]["configType"] {
	case "alioss":
		// 初始化阿里云OSS客户端
		a.AliOSS = internal.NewOssClient(a.ConfigMap["alioss"]["endPoint"], a.ConfigMap["alioss"]["accessKeyId"], a.ConfigMap["alioss"]["accessKeySecret"], a.Log)
	case "limgpath":
		// 配置本地图床文件夹路径
		// a.LocalImgPath = a.ConfigMap["imgbed"]["path"]
	default:
		a.Log.Info(configs.NoImgBedConfigErr)
	}
}

func (a *App) initClipboard() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		a.Log.Error(configs.InitSysClipboardErr, err.Error())
		panic(err)
	}
}

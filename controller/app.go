package gtools

import (
	"context"
	"fmt"
	"gtools/configs"
	"gtools/internal"
	"gtools/util"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

// App 结构体
type App struct {
	ctx          context.Context
	Log          *logrus.Logger
	Db           *xorm.Engine
	LogFile      string
	DBFile       string
	ConfigMap    map[string]map[string]string
	AliOSS       *oss.Client
	LocalImgPath string
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
}

// OnBeforeClose
func (a *App) OnBeforeClose(ctx context.Context) bool {
	// 关闭xorm连接
	a.Db.Close()
	// 返回 true 将阻止程序关闭
	return false
}

// 系统配置数据初始化
func (a *App) initConfigData(ctype string) {
	switch ctype {
	case "alioss":
		datas := make([]*internal.ConfigItem, 0)
		for i := 0; i < len(configs.AliOSS); i++ {
			datas = append(datas, &internal.ConfigItem{Name: configs.AliOSS[i], Value: "", Type: ctype})
		}
		if _, err := a.Db.Insert(&datas); err != nil {
			a.Log.Error(fmt.Sprintf(configs.AddConfigItemErr, ctype, err.Error()))
		}
	case "imgBed":
		datas := make([]*internal.ConfigItem, 0)
		for i := 0; i < len(configs.ImgBed); i++ {
			datas = append(datas, &internal.ConfigItem{Name: configs.ImgBed[i], Value: "", Type: ctype})
		}
		if _, err := a.Db.Insert(&datas); err != nil {
			a.Log.Error(fmt.Sprintf(configs.AddConfigItemErr, ctype, err.Error()))
		}
	case "localImgPath":
		datas := make([]*internal.ConfigItem, 0)
		for i := 0; i < len(configs.LocalImgPath); i++ {
			datas = append(datas, &internal.ConfigItem{Name: configs.LocalImgPath[i], Value: "", Type: ctype})
		}
		if _, err := a.Db.Insert(&datas); err != nil {
			a.Log.Error(fmt.Sprintf(configs.AddConfigItemErr, ctype, err.Error()))
		}
	default:
		a.Log.Warn(fmt.Sprintf("未找到配置项[%s], 系统配置初始化失败", ctype))
	}
}

func (a *App) initImgBed() {
	switch a.ConfigMap["imgBed"]["configType"] {
	case "alioss":
		// 初始化阿里云OSS客户端
		a.AliOSS = internal.NewOssClient(a.ConfigMap["alioss"]["endPoint"], a.ConfigMap["alioss"]["accessKeyId"], a.ConfigMap["alioss"]["accessKeySecret"], a.Log)
	case "localImgPath":
		// 配置本地图床文件夹路径
		a.LocalImgPath = a.ConfigMap["imgBed"]["path"]
	default:
		a.Log.Info(configs.NoImgBedConfigErr)
	}
}

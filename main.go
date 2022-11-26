package main

import (
	"embed"
	"fmt"
	gtools "gtools/controller"
	"net/http"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed frontend/dist
var assets embed.FS

// icon会默认使用 build/appicon.png 转换为byte数组
var icon []byte

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	fileData, err := os.ReadFile("/" + requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}
	res.Write(fileData)
}

func main() {
	// Create an instance of the app structure
	app := gtools.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "GTools", // 标题
		Width:             1100,     // 启动宽度
		Height:            768,      // 启动高度
		MinWidth:          1100,     // 最小宽度
		MinHeight:         768,      // 最小高度
		HideWindowOnClose: true,     // 关闭的时候隐藏窗口
		StartHidden:       true,    // 启动的时候隐藏窗口 （建议生产环境关闭此项，开发环境开启此项，原因自己体会）
		AlwaysOnTop:       false,    // 窗口固定在最顶层
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 128},
		OnStartup:        app.OnStartup,
		OnDomReady:       app.OnDOMReady,
		OnBeforeClose:    app.OnBeforeClose,
		CSSDragProperty:  "--wails-draggable",
		CSSDragValue:     "drag",
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.DefaultAppearance,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "GTools",
				Message: "© 2022 Pixiao",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

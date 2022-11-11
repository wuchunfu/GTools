package main

import (
	"embed"
	gtools "gtools/controller"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"golang.design/x/clipboard"
)

//go:embed frontend/dist
var assets embed.FS

// icon会默认使用 build/appicon.png 转换为byte数组
var icon []byte

func init() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
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
		StartHidden:       false,    // 启动的时候隐藏窗口 （建议生产环境关闭此项，开发环境开启此项，原因自己体会）
		AlwaysOnTop:       false,    // 窗口固定在最顶层
		Assets:            assets,
		BackgroundColour:  &options.RGBA{R: 0, G: 0, B: 0, A: 128},
		OnStartup:         app.OnStartup,
		OnBeforeClose:     app.OnBeforeClose,
		CSSDragProperty:   "--wails-draggable",
		CSSDragValue:      "drag",
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

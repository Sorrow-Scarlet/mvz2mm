package main

import (
	"embed"
	"log/slog"
	"mvz2mm/internal/dashboard"
	"mvz2mm/internal/modmanage"
	"mvz2mm/internal/platform"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/build
var assets embed.FS

func init() {
	// 按照如下格式注册API
	// application.RegisterEvent[string]("time")
}

func init() {
	// TODO: 使用slog替换原有日志
}

func main() {

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "mvz2mm",
		Description: "A MVZ2 Mod Manager",
		Services: []application.Service{
			application.NewService(&platform.PlatformService{}),
			application.NewService(&modmanage.DirService{}),
			application.NewService(&dashboard.StatusService{}),
			application.NewService(&dashboard.VersionService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "MVZ2 Mod Manager",
		// Window sized to the golden ratio (1000 / 618 ≈ 1.618).
		Width:  1000,
		Height: 618,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
		Frameless:        true,
		Windows: application.WindowsWindow{
			NonClientRegionSupport:     true,
			WebView2CompositionHosting: true,
		},
	})

	err := app.Run()
	if err != nil {
		slog.Error("Main Thread", "err", err)
		return
	}
}

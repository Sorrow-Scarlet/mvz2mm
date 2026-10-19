package main

import (
	"embed"
	"io"
	"log/slog"
	"mvz2mm/internal/modmanage"
	"mvz2mm/internal/platform"
	"mvz2mm/internal/settings"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/build
var assets embed.FS

func init() {
	isProd := !strings.Contains(settings.AppVersion, "-") &&
		strings.Count(settings.AppVersion, ".") == 2

	dir := "./logs"
	_ = os.MkdirAll(dir, 0o755)

	f, err := os.OpenFile(
		filepath.Join(dir, "mvz2mm-"+time.Now().Format("20060102150405")+".log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)

	if err != nil {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
		return
	}

	if isProd {
		slog.SetDefault(slog.New(slog.NewJSONHandler(f, nil)))
	} else {
		slog.SetDefault(slog.New(slog.NewTextHandler(
			io.MultiWriter(os.Stdout, f), nil)))
	}
}

func main() {

	dirService := &settings.DirService{}
	dirService.CheckGameDir()

	app := application.New(application.Options{
		Name:        "mvz2mm",
		Description: "A MVZ2 Mod Manager",
		Services: []application.Service{
			application.NewService(&platform.PlatformService{}),
			application.NewService(&modmanage.ModService{}),
			application.NewService(dirService),
			application.NewService(&settings.VersionService{}),
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
		URL:              "/dashboard",
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

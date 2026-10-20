package settings

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v3/pkg/application"
)

const (
	ExportmodDllDir   = modDllDir
	Exportmvz2DataDir = mvz2DataDir
)

const (
	mvz2DataDir    = "MinecraftVSZombies2_Data"
	modDllDir      = "BepInEx/plugins"
	modResourceDir = mvz2DataDir + "/StreamingAssets/Mods"
)

type DirService struct {
	GameDir string `json:"gameDir"`
}

func NewDirService() *DirService {
	return &DirService{}
}

// region 前端API

// GetGameDir 仅返回MVZ2**根目录**，未设置时为空。
func (d *DirService) GetGameDir() string {
	return d.GameDir
}

func (d *DirService) GetModDllDir() string {
	return d.joinGameDir(modDllDir)
}

// GetModResourceDir 返回mod资源目录的绝对路径
func (d *DirService) GetModResourceDir() string {
	return d.joinGameDir(modResourceDir)
}
func (d *DirService) SelectGameDir() error {
	exePath, err := application.Get().Dialog.OpenFile().
		SetTitle("选择MVZ2游戏主程序").
		PromptForSingleSelection()
	if err != nil {
		slog.Error("select mvz2 exe failed", "err", err)
		return errors.New("打开文件选择器失败: " + err.Error())
	}
	if exePath == "" {
		return nil // 用户取消
	}

	gameDir := resolveGameDir(exePath)
	if gameDir == "" {
		slog.Warn("selected exe has no data folder beside it", "exePath", exePath)
		return errors.New("所选主程序旁未找到 " + mvz2DataDir + " 文件夹，请确认选择的是MVZ2游戏目录下的主程序")
	}

	d.GameDir = gameDir
	slog.Info("game dir set via selection", "gameDir", d.GameDir)
	return nil
}

func (d *DirService) CheckGameDir() {
	exePath, err := os.Executable()
	if err != nil {
		slog.Error("get executable path failed", "err", err)
		return
	}
	appDir := filepath.Dir(exePath)
	if isGameDir(appDir) {
		d.GameDir = appDir
		slog.Info("game dir auto detected", "gameDir", d.GameDir)
		return
	}
	slog.Info("game dir not detected beside executable", "appDir", appDir)
}

// region ModService预留

// EnsureModResourceDir 确保mod资源目录存在，不存在则创建。
func (d *DirService) EnsureModResourceDir() error {
	if d.GameDir == "" {
		return errors.New("游戏目录未设置，无法创建mod资源目录")
	}
	if err := os.MkdirAll(d.GetModResourceDir(), 0o755); err != nil {
		slog.Error("create mod resource dir failed", "dir", d.GetModResourceDir(), "err", err)
		return errors.New("创建mod资源目录失败: " + err.Error())
	}
	return nil
}
func (d *DirService) HasGameDirSet() bool {
	return d.GameDir != ""
}

// region PRIVATE

// joinGameDir 拼接游戏根目录与相对路径，未设置游戏目录时返回空串。
func (d *DirService) joinGameDir(rel string) string {
	if d.GameDir == "" {
		return ""
	}
	return filepath.Join(d.GameDir, filepath.FromSlash(rel))
}

// resolveGameDir 由主程序exe路径推导游戏根目录，不是合法游戏目录时返回空串。
func resolveGameDir(exePath string) string {
	gameDir := filepath.Dir(exePath)
	if isGameDir(gameDir) {
		return gameDir
	}
	return ""
}

// isGameDir 判断目录下是否存在MVZ2数据文件夹。
func isGameDir(dir string) bool {
	info, err := os.Stat(filepath.Join(dir, mvz2DataDir))
	if err != nil {
		return false
	}
	return info.IsDir()
}

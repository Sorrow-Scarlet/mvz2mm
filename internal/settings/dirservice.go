package settings

import (
	"errors"
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type DirService struct {
	GameDir string `json:"game_dir"`
	ModDir  string `json:"mod_dir"`
}

const mvz2ExeName = ""
const mvz2DataFolder = "MinecraftVSZombies2_Data"

func (d *DirService) GetGameDir() string {
	return d.GameDir
}

func (d *DirService) SetGameDir(path string) error {
	if !isGameDir(path, mvz2DataFolder) {
		return errors.New("找不到MVZ2数据文件夹")
	}
	d.GameDir = path
	return nil
}

// CheckGameDir() 检查根目录是否为MVZ2目录，main.go初始化用
func (d *DirService) CheckGameDir() {

	currentDir, err := os.Getwd()
	if err != nil {
		slog.Error("get current app dir", "err", err)
		return
	}
	if isGameDir(currentDir, mvz2DataFolder) {
		d.GameDir = currentDir
		slog.Info("set game dir", "current d.GameDir", d.GameDir)
		return
	}

}

func (d *DirService) GetGameDirFromDialogSelection(pathWithMVZ2 string) string {
	p, _ := path.Split(pathWithMVZ2)
	if !isGameDir(mvz2DataFolder, p) {
		return ""
	}
	return p
}

func (d *DirService) GetModDir() string {
	if d.GameDir == "" {
		slog.Warn("gameDir is empty!")
		return ""
	}

	tryPath := path.Join(d.GameDir, "BepInEx", "plugins")
	info, err := os.Stat(tryPath)
	if err != nil {
		slog.Warn("ModDir dose not exists!", "err", err, "dir", tryPath)
		return ""
	}
	if info.IsDir() {
		d.ModDir = tryPath
		return d.ModDir
	}
	return ""
}

// SelectGameDir() 处理前端`设置页MVZ2路径选择按钮`的事件
func (d *DirService) SelectGameDir() error {
	path, err := application.Get().Dialog.OpenFile().SetTitle("选择MVZ2游戏主程序").PromptForSingleSelection()
	if err != nil || path == "" {
		slog.Error("selecting mvz2 game", "err", err, "path", path)
		return errors.New("未找到主程序或用户取消选择")
	}

	pS := d.GetGameDirFromDialogSelection(path)
	err = d.SetGameDir(pS)
	if err != nil {
		slog.Error("setting game dir via selection", "err", err)
		return err
	}

	slog.Info("set game file", "current d.GameDir", d.GameDir)
	return nil
}

// isGameDir 接收dataFolder和已有路径为参数，确认是否为MVZ2目录
//
// 若传入路径 **包含** dataFolder则返回真
func isGameDir(dataFolder, path string) bool {
	info, err := os.Stat(filepath.Join(path, dataFolder))
	// 没出错，存在这个文件，
	if err == nil {

	}
	// 出错，不存在此文件/目录，不是MVZ2目录
	if err != nil {
		slog.Warn("scanning game data folder", "err", err)
		return false
	}
	return info.IsDir()
}

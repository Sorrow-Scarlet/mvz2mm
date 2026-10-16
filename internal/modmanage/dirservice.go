package modmanage

import (
	"log/slog"
	"os"
	"path"
	"path/filepath"
)

type DirService struct {
	GameDir string `json:"game_dir"`
	ModDir  string `json:"mod_dir"`
}

func (d *DirService) GetGameDir() string {

	const mvz2DataFolder = "MinecraftVSZombies2_Data"

	currentDir, err := os.Getwd()
	if err != nil {
		slog.Error("get current app dir", "err", err)
		return ""
	}
	if isGameDir(currentDir, mvz2DataFolder) {
		d.GameDir = currentDir
		return d.GameDir
	}

	return ""
}

func (d *DirService) GetModDir(gameDir string) string {
	if gameDir == "" {
		slog.Warn("gameDir is empty!")
		return ""
	}

	tryPath := path.Join(gameDir, "BepInEx", "plugins")
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

// isGameDir 接收dataFolder和已有路径为参数，确认是否为MVZ2目录
//
// 若传入路径 **包含** dataFolder则返回真
func isGameDir(dataFolder, path string) bool {
	info, err := os.Stat(filepath.Join(path, dataFolder))
	// 出错，不存在此目录，不是MVZ2目录
	if err != nil {
		slog.Warn("scanning game data folder", "err", err)
		return false
	}
	return info.IsDir()
}

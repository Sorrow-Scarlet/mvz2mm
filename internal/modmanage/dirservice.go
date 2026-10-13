package modmanage

import (
	"log/slog"
	"os"
	"path/filepath"
)

type DirService struct{}

func (d *DirService) GetApplicationDir() (string, error) {
	return os.Getwd()
}

func (d *DirService) JoinMVZ2Path(dir, MVZ2ExeName string) string {
	return filepath.Join(dir, MVZ2ExeName)
}

func (d *DirService) IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		slog.Info("checking path", "err", err)
		return false
	}
	return true
}

func (d *DirService) ChooseGameDir() {}

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

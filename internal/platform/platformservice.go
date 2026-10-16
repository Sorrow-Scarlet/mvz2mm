package platform

import (
	"mvz2mm/internal/platform/gamemeta"
	"mvz2mm/internal/platform/opendir"
)

type PlatformService struct{}

// OpenDir 传入绝对路径，使用`生字符串`，启动系统资源管理器
func (p *PlatformService) OpenDir(dir string) error {
	return opendir.Open(dir)
}

func (p *PlatformService) ReadMeta(exePath string) (gamemeta.GameMetaInfo, error) {
	return gamemeta.Read(exePath)
}

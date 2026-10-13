package platform

import (
	"mvz2mm/internal/platform/opendir"
	"mvz2mm/internal/platform/readmeta"
)

type PlatformService struct{}

// OpenDir 传入绝对路径，使用`生字符串`，启动系统资源管理器
func (p *PlatformService) OpenDir(dir string) error {
	return opendir.Open(dir)
}

func (p *PlatformService) ReadMeta(exePath string) (readmeta.GameMetaInfo, error) {
	return readmeta.Read(exePath)
}

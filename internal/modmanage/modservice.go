package modmanage

import "mvz2mm/internal/settings"

// RawManifest 融合了标准单MOD和合集MOD，其复用`名称`与作者
type RawManifest struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Author       string   `json:"author"`
	Dependencies []string `json:"dependencies"`
	HasResources bool     `json:"has_resources"`
	// 合集专有信息
	CollectionMod []string `json:"Mods"`
	CreatedUtc    string   `json:"CreatedUtc"`
}

type Mod struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
	Enabled bool   `json:"enabled"`
}

var GlobalModList = make([]Mod, 5)

type ModService struct {
	DS         settings.DirService
	ModTotal   int
	ModEnabled int
}

func (m *ModService) InitModStatus() error {

	return nil
}

func (m *ModService) GetModTotal() int {
	return m.ModTotal
}

func (m *ModService) GetModEnabled() int {
	return m.ModEnabled
}

func ChangeModEnable() {}

func InstallMod(mod Mod) {}

func UninstallMod(mod Mod) {}

func (m *ModService) scanDLLFiles() {
	const extname = "dll"
}

func (m *ModService) scanDisabled() {}

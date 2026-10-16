package modmanage

import "mvz2mm/internal/settings"

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

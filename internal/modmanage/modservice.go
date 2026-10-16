package modmanage

import "mvz2mm/internal/settings"

type ModService struct {
	DS settings.DirService
}

func (m *ModService) GetModTotal() (int, error) {

	return 1, nil
}

func (m *ModService) GetModEnabled() (int, error) {
	return 0, nil
}

func (m *ModService) ScanDLLFiles(path string) {

}

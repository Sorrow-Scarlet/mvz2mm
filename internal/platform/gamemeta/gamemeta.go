package gamemeta

import (
	"log/slog"
	"runtime"
)

type GameMetaInfo struct {
	ExecutableFileVersion    string `json:"file_version"`
	ExecutableProductVersion string `json:"product_version"`
}

func Read(exePath string) (GameMetaInfo, error) {
	info, err := readMeta(exePath)
	if err != nil {
		slog.Error("reading mvz2 meta info", "err", err, "OS", runtime.GOOS, "path", exePath)
		return GameMetaInfo{}, err
	}
	return info, nil
}

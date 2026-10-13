package opendir

import (
	"log/slog"
	"runtime"
)

func Open(dir string) error {
	err := openDir(dir)
	if err != nil {
		slog.Error("opening system file explorer", "err", err, "dir", dir, "os", runtime.GOOS)
		return err
	}
	slog.Debug("opened system explorer")
	return nil
}

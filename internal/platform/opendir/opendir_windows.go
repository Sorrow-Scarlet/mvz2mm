//go:build windows

package opendir

import "os/exec"

func openDir(dir string) error {
	return exec.Command(`cmd`, `/c`, `explorer`, dir).Start()
}

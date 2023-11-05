package util

import (
	"os/exec"
	"runtime"
)

func OpenLoggoDebug(file string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open -a Terminal"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, "loggo", "stream", "--file", file)
	return exec.Command(cmd, args...).Start()
}

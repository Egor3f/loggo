package util

import (
	"os/exec"
	"runtime"
	"strings"
)

func OpenBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	Log().WithField("code", cmd+" "+strings.Join(args, " ")).Info("Issue browser command.")
	return exec.Command(cmd, args...).Start()
}

package main

import (
	"os/exec"
	"runtime"
)

func openURL(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default: // for linux and freebsd
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}

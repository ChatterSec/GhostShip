package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

func exit() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()

	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	os.Exit(0)
}

func main() {
	// listen for Ctrl+C, and when it's received, revert the terminal settings
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		exit()
	}()

	if runtime.GOOS == "windows" {
		setConsoleMode()
	} else if runtime.GOOS == "darwin" {
		exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run() // disable input buffering
		exec.Command("stty", "-f", "/dev/tty", "-echo").Run()              // do not display entered characters on the screen
	} else if runtime.GOOS == "linux" {
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run() // disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()              // do not display entered characters on the screen
	}

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		fmt.Println(b)
	}
}

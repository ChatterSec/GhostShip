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

	hovering := 0
	options := []string{"module", "booty", "help", "report", "exit"}
	colors := make([]string, len(options))
	cursor := make([]string, len(options))
	first_loop := true

	var b []byte = make([]byte, 1)
	for {
		if !first_loop {
			os.Stdin.Read(b)
		}

		// Up
		if b[0] == 107 {
			if hovering == 0 {
				hovering = len(options) - 1
			} else {
				hovering -= 1
			}
		}

		// Down
		if b[0] == 106 {
			if hovering == len(options)-1 {
				hovering = 0
			} else {
				hovering += 1
			}
		}

		if b[0] == 106 || b[0] == 107 || first_loop {
			for i := range options {
				if i == hovering {
					colors[i] = "\033[32m"
					cursor[i] = ">"
				} else {
					colors[i] = "\033[0m"
					cursor[i] = "-"
				}
			}

			if !first_loop {
				fmt.Printf("\033[%dA", 16)
				fmt.Print("\033[2K")
			}

			fmt.Println("\033[32m\033[1m" + `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣴⣶⣤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣠⡤⣤⣄⣾⣿⣿⣿⣿⣿⣿⣷⣠⣀⣄⡀` + "\033[0m\033[1m\t " + `GhostShip v0.1.0` + "\033[32m\033[1m" + `
⠀⠀⠀⠀⠙⠀⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣬⡿` + "\033[0m\t by: \033]8;;https://reeceharris.net\033\\notreeceharris\033]8;;\033\\\033[32m\033[1m" + `
⠀⠀⠀⠀⠀⢀⣼⠟⢿⣿⣿⣿⣿⣿⣿⡿⠘⣷⣄
⣰⠛⠛⣿⢠⣿⠋⠀⠀⢹⠻⣿⣿⡿⢻⠁⠀⠈⢿⣦` + "\033[0m\033[2m\t " + `A pirate-themed penetration testing framework,` + "\033[0m\033[32m\033[1m" + `
⢈⣵⡾⠋⣿⣯⠀⠀⢀⣼⣷⣿⣿⣶⣷⡀⠀⠀⢸⣿⣀⣀` + "\033[0m\033[2m\t " + `built using reliable tools and custom scripts.` + "\033[0m\033[32m\033[1m" + `
⢾⣿⣀⠀⠘⠻⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⠿⣿⡁
⠈⠙⠛⠿⠿⠿⢿⣿⡿⣿⣿⡿⢿⣿⣿⣿⣷⣄⠀⠘⢷⣆` + "\033[0m" + colors[0] + "\t " + cursor[0] + ` modules` + "\033[32m\033[1m" + `
⠀⠀⠀⠀⠀⢠⣿⠏⠀⣿⡏⠀⣼⣿⠛⢿⣿⣿⣆⠀⠀⣿⡇⡀` + "\033[0m" + colors[1] + " " + cursor[1] + " booty \033[2m(captures)" + "\033[0m\033[32m\033[1m" + `⠀
⠀⠀⠀⠀⢀⣾⡟⠀⠀⣿⣇⠀⢿⣿⡀⠈⣿⡌⠻⠷⠾⠿⣻⠁` + "\033[0m" + colors[2] + " " + cursor[2] + " sos \033[2m(help)" + "\033[0m\033[32m\033[1m" + `⠀
⠀⠀⣠⣶⠟⠫⣤⠀⠀⢸⣿⠀⣸⣿⢇⡤⢼⣧⠀⠀⠀⢀⣿` + "\033[0m" + colors[3] + "\t " + cursor[3] + ` report issue` + "\033[32m\033[1m" + `⠀
⠀⣾⡏⠀⡀⣠⡟⠀⠀⢀⣿⣾⠟⠁⣿⡄⠀⠻⣷⣤⣤⡾⠋` + "\033[0m" + colors[4] + "\t " + cursor[4] + " disembark \033[2m(exit)" + "\033[0m\033[32m\033[1m" + `⠀
⠀⠙⠷⠾⠁⠻⣧⣀⣤⣾⣿⠋⠀⠀⢸⣧⠀⠀⠀⠉⠁
⠀⠀⠀⠀⠀⠀⠈⠉⠉⠹⣿⣄⠀⠀⣸⡿` + "\033[0m\033[2m\t " + `(k = up, j = down, enter = submit)` + "\033[0m\033[32m\033[1m" + `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠛⠿⠟⠛⠁` + "\033[0m")
		}

		first_loop = false
	}
}

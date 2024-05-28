package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func exit() {

	// Revert the terminal settings before exiting
	if runtime.GOOS == "windows" {
		resetConsoleMode()
	} else if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		var flag string
		if runtime.GOOS == "darwin" {
			flag = "-f"
		} else {
			flag = "-F"
		}
		exec.Command("stty", flag, "/dev/tty", "icanon", "echo").Run()
	}

	os.Exit(0)
}

func init() {
	// listen for Ctrl+C, and when it's received, revert the terminal settings
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		exit()
	}()

	// Disable inputs
	if runtime.GOOS == "windows" {
		setConsoleMode()
	} else if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		var flag string
		if runtime.GOOS == "darwin" {
			flag = "-f"
		} else {
			flag = "-F"
		}
		exec.Command("stty", flag, "/dev/tty", "cbreak", "min", "1").Run() // disable input buffering
		exec.Command("stty", flag, "/dev/tty", "-echo").Run()              // do not display entered characters on the screen
	}
}

func main() {

	var (
		hovering          = 0
		options           = []string{"module", "booty", "help", "report", "exit"}
		colors            = make([]string, len(options))
		cursor            = make([]string, len(options))
		first_loop        = true
		b          []byte = make([]byte, 1)
	)

	ke := "\033[0m\033[32m\033[1m"
	kraken := []string{
		"⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣴⣶⣤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀",
		"⠀⠀⠀⠀⣠⡤⣤⣄⣾⣿⣿⣿⣿⣿⣿⣷⣠⣀⣄⡀⠀⠀⠀⠀",
		"⠀⠀⠀⠀⠙⠀⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣬⡿⠀⠀⠀⠀",
		"⠀⠀⠀⠀⠀⢀⣼⠟⢿⣿⣿⣿⣿⣿⣿⡿⠘⣷⣄⠀⠀⠀⠀⠀",
		"⣰⠛⠛⣿⢠⣿⠋⠀⠀⢹⠻⣿⣿⡿⢻⠁⠀⠈⢿⣦⠀⠀⠀⠀",
		"⢈⣵⡾⠋⣿⣯⠀⠀⢀⣼⣷⣿⣿⣶⣷⡀⠀⠀⢸⣿⣀⣀⠀⠀",
		"⢾⣿⣀⠀⠘⠻⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⠿⣿⡁⠀⠀⠀",
		"⠈⠙⠛⠿⠿⠿⢿⣿⡿⣿⣿⡿⢿⣿⣿⣿⣷⣄⠀⠘⢷⣆⠀⠀",
		"⠀⠀⠀⠀⠀⢠⣿⠏⠀⣿⡏⠀⣼⣿⠛⢿⣿⣿⣆⠀⠀⣿⡇⡀",
		"⠀⠀⠀⠀⢀⣾⡟⠀⠀⣿⣇⠀⢿⣿⡀⠈⣿⡌⠻⠷⠾⠿⣻⠁",
		"⠀⠀⣠⣶⠟⠫⣤⠀⠀⢸⣿⠀⣸⣿⢇⡤⢼⣧⠀⠀⠀⢀⣿⠀",
		"⠀⣾⡏⠀⡀⣠⡟⠀⠀⢀⣿⣾⠟⠁⣿⡄⠀⠻⣷⣤⣤⡾⠋⠀",
		"⠀⠙⠷⠾⠁⠻⣧⣀⣤⣾⣿⠋⠀⠀⢸⣧⠀⠀⠀⠉⠁⠀⠀⠀",
		"⠀⠀⠀⠀⠀⠀⠈⠉⠉⠹⣿⣄⠀⠀⣸⡿⠀⠀⠀⠀⠀⠀⠀⠀",
		"⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠛⠿⠟⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀",
	}

	// Put random spots to the kraken
	rand.Seed(time.Now().Unix())
	for i, val := range kraken {
		changed := 0
		chars := []rune(val)
		for changed < 30 {
			index := rand.Intn(len(chars))
			if chars[index] != 10240 {
				newStr := string(chars[:index]) + "\x1b[2m" + string(chars[index]) + "\033[0m\033[32m\033[1m" + string(chars[index+1:])
				kraken[i] = newStr
				changed++
			}
		}
	}

	for {
		if !first_loop {
			os.Stdin.Read(b)
		}

		// Handle "k" Key Press (aka up)
		if b[0] == 107 {
			if hovering == 0 {
				hovering = len(options) - 1
			} else {
				hovering -= 1
			}
		}

		// Handle "j" Key Press (aka down)
		if b[0] == 106 {
			if hovering == len(options)-1 {
				hovering = 0
			} else {
				hovering += 1
			}
		}

		// Handle Space Bar Press
		if b[0] == 32 {
			if hovering == 3 { // Report an issue.
				exit()
			}

			if hovering == 4 { // Exit
				exit()
			}
		}

		if b[0] == 106 || b[0] == 107 || first_loop {
			for i := range options {
				if i == hovering {
					colors[i] = "\033[32m"
					cursor[i] = " >"
				} else {
					colors[i] = "\033[0m"
					cursor[i] = " -"
				}
			}

			if !first_loop {
				fmt.Printf("\033[%dA", 17)
				fmt.Print("\033[2K")
			}

			// Print Menu
			fmt.Println("")
			fmt.Println(ke + kraken[0] + "\033[0m")
			fmt.Println(ke + kraken[1] + "\033[0m\033[1m GhostShip v0.1.0")
			fmt.Println(ke + kraken[2] + "\033[0m by: \033]8;;https://reeceharris.net\033\\notreeceharris\033]8;;\033\\\033[32m\033[1m")
			fmt.Println(ke + kraken[3] + "\033[0m")
			fmt.Println(ke + kraken[4] + "\033[0m\033[2m A pirate-themed penetration testing framework,")
			fmt.Println(ke + kraken[5] + "\033[0m\033[2m built using reliable tools and custom scripts.")
			fmt.Println(ke + kraken[6] + "\033[0m")
			fmt.Println(ke + kraken[7] + "\033[0m" + colors[0] + cursor[0] + " man-o-war    \033[2m(tools)")
			fmt.Println(ke + kraken[8] + "\033[0m" + colors[1] + cursor[1] + " booty        \033[2m(captures)")
			fmt.Println(ke + kraken[9] + "\033[0m" + colors[2] + cursor[2] + " sos          \033[2m(help)")
			fmt.Println(ke + kraken[10] + "\033[0m" + colors[3] + cursor[3] + " bounty       \033[2m(report issue)")
			fmt.Println(ke + kraken[11] + "\033[0m" + colors[4] + cursor[4] + " disembark    \033[2m(exit)")
			fmt.Println(ke + kraken[12] + "\033[0m")
			fmt.Println(ke + kraken[13] + "\033[0m\033[2m" + ` (k = up, j = down, space = submit)`)
			fmt.Println(ke + kraken[14] + "\033[0m")
			fmt.Println("")
		}

		first_loop = false
	}
}

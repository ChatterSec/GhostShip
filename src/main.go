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

func clearTerminal() {
	fmt.Printf("\033[%dA", 17)
	fmt.Print("\033[2K")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
	fmt.Println("                                                                                    ")
}

func toolsMenu() {

	var (
		hovering          = 0
		options           = []string{"all", "osint", "exploitation", "post-exploitation", "reporting", "settings", "search", "back"}
		colors            = make([]string, len(options))
		cursor            = make([]string, len(options))
		first_loop        = true
		b          []byte = make([]byte, 1)
	)

	cutlass := []string{
		"   \x1b[31m^\033[0m     ",
		"  \x1b[31m|\033[0m.|    ",
		"  |\x1b[31m|\033[0m|    ",
		"  |\x1b[31m|\033[0m|    ",
		"  \x1b[31m|\033[0m||    ",
		"  \x1b[31m|\033[0m||    ",
		"  |\x1b[31m|\033[0m|    ",
		"\033[33m__\033[0m||\x1b[31m|\033[0m\033[33m_   \033[0m",
		"\033[33m`----.`  \033[0m",
		"\033[33m  ||  )) \033[0m",
		"\033[33m  |'-',  \033[0m",
		"\033[33m   '-'   \033[0m",
	}

	for {
		if !first_loop {
			os.Stdin.Read(b)
		}

		// Handle "k" Key Press (aka up)
		if b[0] == 107 || b[0] == 75 {
			if hovering == 0 {
				hovering = len(options) - 1
			} else {
				hovering -= 1
			}
		}

		// Handle "j" Key Press (aka down)
		if b[0] == 106 || b[0] == 74 {
			if hovering == len(options)-1 {
				hovering = 0
			} else {
				hovering += 1
			}
		}

		// Handle Space Bar Press
		if b[0] == 32 {

			if hovering == 7 { // Back
				clearTerminal()
				mainMenu(true)
			}
		}

		if b[0] == 106 || b[0] == 74 || b[0] == 107 || b[0] == 75 || first_loop {
			for i := range options {
				if i == hovering {
					colors[i] = "\033[33m"
					cursor[i] = " >"
				} else {
					colors[i] = "\033[0m"
					cursor[i] = " -"
				}
			}

			fmt.Printf("\033[%dA", 17)
			fmt.Print("\033[2K")

			// Print Menu

			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println(cutlass[0])
			fmt.Println(cutlass[1] + colors[0] + cursor[0] + " All   \033[2m(58 Tools)" + "\033[0m")
			fmt.Println(cutlass[2] + colors[1] + cursor[1] + " Osint \033[2m(Open-source Intelligence)" + "\033[0m")
			fmt.Println(cutlass[3] + colors[2] + cursor[2] + " Exploitation" + "\033[0m")
			fmt.Println(cutlass[4] + colors[3] + cursor[3] + " Post-Exploitation" + "\033[0m")
			fmt.Println(cutlass[5] + colors[4] + cursor[4] + " Reporting" + "\033[0m")
			fmt.Println(cutlass[6])
			fmt.Println(cutlass[7])
			fmt.Println(cutlass[8] + colors[5] + cursor[5] + " Settings" + "\033[0m")
			fmt.Println(cutlass[9] + colors[6] + cursor[6] + " Search" + "\033[0m")
			fmt.Println(cutlass[10] + colors[7] + cursor[7] + " Go Back" + "\033[0m")
			fmt.Println(cutlass[11])
			fmt.Println(" ")
			fmt.Println("\033[2m (k = up, j = down, space = submit) \033[0m")
			fmt.Println(" ")

			/* fmt.Println(" ")
			fmt.Println(colors[0] + cursor[0] + " Go Back" + "\033[0m")
			fmt.Println(" ")
			fmt.Println("  OSINT TOOLS       |  EXPLOITATION TOOLS  |  POST-EXPLOITATION TOOLS")
			fmt.Println(" -------------------|----------------------|-------------------------")
			fmt.Println(colors[1] + cursor[1] + " DetectDee" + "\033[0m        | " + colors[2] + cursor[2] + " SqlMap" + "\033[0m            | " + colors[3] + cursor[3] + " Reporting")
			fmt.Println(colors[1] + cursor[1] + " WhatWeb" + "\033[0m          | " + colors[2] + cursor[2] + " Netcat" + "\033[0m            | " + colors[3] + cursor[3] + " Ghost Framework")
			fmt.Println(colors[1] + cursor[1] + " Nmap" + "\033[0m             | " + colors[2] + cursor[2] + " Aircrack-ng" + "\033[0m       | " + colors[3] + cursor[3] + " SSH-Snake")
			fmt.Println(colors[1] + cursor[1] + " Recon-ng" + "\033[0m         | " + colors[2] + cursor[2] + " Hashcat" + "\033[0m           | " + colors[3] + cursor[3] + " psEmpire")
			fmt.Println(colors[1] + cursor[1] + " theHarvester" + "\033[0m     | " + colors[2] + cursor[2] + " Hydra" + "\033[0m             | " + colors[3] + cursor[3] + " Mimikatz")
			fmt.Println(colors[1] + cursor[1] + "     " + "\033[0m             | " + colors[2] + cursor[2] + " Nikto" + "\033[0m             | " + colors[3] + cursor[3] + " LaZagne")
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println(" ") */
		}

		first_loop = false
	}
}

func mainMenu(init_clear bool) {

	var (
		hovering          = 0
		options           = []string{"module", "booty", "help", "report", "exit"}
		colors            = make([]string, len(options))
		cursor            = make([]string, len(options))
		first_loop        = true
		clear             = init_clear
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
		tries := 0
		changedArr := []int{}
		chars := []rune(val)
		for changed < 20 && tries < 50 {
			index := rand.Intn(len(chars))

			if chars[index] != 10240 {
				found := false
				for _, v := range changedArr {
					if v == index {
						found = true
						break
					}
				}
				if !found {
					changedArr = append(changedArr, index)
					newStr := string(chars[:index]) + "\x1b[2m" + string(chars[index]) + "\033[0m" + "\033[32m\033[1m" + string(chars[index+1:])
					kraken[i] = newStr
					changed++
				}
			}
			tries++
		}
	}

	for {
		if !first_loop {
			os.Stdin.Read(b)
		}

		// Handle "k" Key Press (aka up)
		if b[0] == 107 || b[0] == 75 {
			if hovering == 0 {
				hovering = len(options) - 1
			} else {
				hovering -= 1
			}
		}

		// Handle "j" Key Press (aka down)
		if b[0] == 106 || b[0] == 74 {
			if hovering == len(options)-1 {
				hovering = 0
			} else {
				hovering += 1
			}
		}

		// Handle Space Bar Press
		if b[0] == 32 {

			if hovering == 0 { // tools
				clearTerminal()
				toolsMenu()
			}

			if hovering == 3 { // Report an issue
				exit()
			}

			if hovering == 4 { // Exit
				exit()
			}
		}

		if b[0] == 106 || b[0] == 74 || b[0] == 107 || b[0] == 75 || first_loop {
			for i := range options {
				if i == hovering {
					colors[i] = "\033[32m"
					cursor[i] = " >"
				} else {
					colors[i] = "\033[0m"
					cursor[i] = " -"
				}
			}

			if !first_loop || clear {
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
			fmt.Println(ke + kraken[7] + "\033[0m" + colors[0] + cursor[0] + " armory       \033[2m(tools)")
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
		clear = false
	}
}

func main() {

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

	mainMenu(false)
}

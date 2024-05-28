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
		options           = []string{"Back", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test", "Test"}
		colors            = make([]string, len(options))
		cursor            = make([]string, len(options))
		first_loop        = true
		b          []byte = make([]byte, 1)
	)

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

			if hovering == 0 { // Back
				clearTerminal()
				mainMenu(true)
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

			fmt.Printf("\033[%dA", 17)
			fmt.Print("\033[2K")

			// Print Menu
			fmt.Println(" ")
			fmt.Println(colors[0] + cursor[0] + " Go Back  " + "\033[0m   | " + colors[15] + cursor[15] + " DetectDee")
			fmt.Println(colors[1] + cursor[1] + " DetectDee" + "\033[0m   | " + colors[16] + cursor[16] + " DetectDee")
			fmt.Println(colors[2] + cursor[2] + " DetectDee" + "\033[0m   | " + colors[17] + cursor[17] + " DetectDee")
			fmt.Println(colors[3] + cursor[3] + " DetectDee" + "\033[0m   | " + colors[18] + cursor[18] + " DetectDee")
			fmt.Println(colors[4] + cursor[4] + " DetectDee" + "\033[0m   | " + colors[19] + cursor[19] + " DetectDee")
			fmt.Println(colors[5] + cursor[5] + " DetectDee" + "\033[0m   | " + colors[20] + cursor[20] + " DetectDee")
			fmt.Println(colors[6] + cursor[6] + " DetectDee" + "\033[0m   | " + colors[21] + cursor[21] + " DetectDee")
			fmt.Println(colors[7] + cursor[7] + " DetectDee" + "\033[0m   | " + colors[22] + cursor[22] + " DetectDee")
			fmt.Println(colors[8] + cursor[8] + " DetectDee" + "\033[0m   | " + colors[23] + cursor[23] + " DetectDee")
			fmt.Println(colors[9] + cursor[9] + " DetectDee" + "\033[0m   | " + colors[24] + cursor[24] + " DetectDee")
			fmt.Println(colors[10] + cursor[10] + " DetectDee" + "\033[0m   | " + colors[25] + cursor[25] + " DetectDee")
			fmt.Println(colors[11] + cursor[11] + " DetectDee" + "\033[0m   | " + colors[26] + cursor[26] + " DetectDee")
			fmt.Println(colors[12] + cursor[12] + " DetectDee" + "\033[0m   | " + colors[27] + cursor[27] + " DetectDee")
			fmt.Println(colors[13] + cursor[13] + " DetectDee" + "\033[0m   | " + colors[28] + cursor[28] + " DetectDee")
			fmt.Println(colors[14] + cursor[14] + " DetectDee" + "\033[0m   | " + colors[29] + cursor[29] + " DetectDee")
			fmt.Println("")
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

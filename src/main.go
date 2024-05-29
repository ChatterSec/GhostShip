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

	"golang.org/x/crypto/ssh/terminal"
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
	for i := 0; i < 17; i++ {
		fmt.Println("                                                                                    ")
	}
}

func toolsMenu() {

	var (
		hovering          = 0
		options           = []string{"all", "osint", "exploitation", "post-exploitation", "reporting", "search", "back"}
		colors            = make([]string, len(options))
		cursor            = make([]string, len(options))
		first_loop        = true
		b          []byte = make([]byte, 1)
	)

	manowar := []string{
		"\033[32m\033[1m⠀⠀⠀⠀⠀⠀⠀⠀⠀\033[0m\033[1m⣀⠀⠤⠴⣶\033[32m⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀\033[0m ",
		"\033[0m\033[1m⠀⠀⠀⠀⠀⠀⠀⠀⠀ ⠉ ⠛⠟\033[32m⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀\033[0m  ",
		"\033[0m\033[1m⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⣾⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀ \033[0m ",
		"\033[32m\033[1m⠀⠀⠀\033[0m\033[1m       \033[0m\033[1m⢰⣿⣿⣿⣿⣧⠀⠀⢀⣄⣀⠀⠀⠀\033[0m ",
		"\033[0m\033[1m⠀⠀⠀⢠⣶⣶⣷⠀⠀⠀⠸⠟⠁⠀\033[32m⡇⠀⠀⠀⠀⠀⢹⠀⠀⠀\033[0m ",
		"\033[0m\033[1m⠀⠀⠀⠘⠟\033[32m⢸\033[0m\033[1m⣋⣀⡀⢀⣤⣶⣿⣿⣿⣿⣿⡿⠛⣠⣼⣿⡟⠀\033[0m ",
		"\033[0m\033[1m⠀⠀⣴⣾⣿⣿⣿⣿⢁⣾⣿⣿⣿⣿⣿⣿⡿⢁⣾⣿⣿⣿⠁⠀\033[0m ",
		"\033[0m\033[1m⠀⠸⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⡇⢸⣿⣿⣿⠿⠇⠀\033[0m ",
		"\033[32m\033[1m⠳⣤⣀\033[0m\033[1m⠘⠛⢻⠿⣿⠸⣿⣿⣿⣿⣿⣿⣿⣇⠘⠉⠀\033[32m⢸⠀⢀⣠\033[0m ",
		"\033[32m\033[1m⠀⠈⠻⣷⣦⣼\033[0m\033[1m⠀⠀⠀⢻⣿⣿⠿⢿⡿⠿⣿⡄⠀⠀\033[32m ⣼⣷⣿⣿\033[0m",
		"\033[32m\033[1m⠀⠀⠀⠈⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⣶⣄\033[0m⡈⠉\033[32m⠀⠀⢸⡇⠀⠀\033[0m\033[1m⠉⠂\033[32m⠀⣿⣿⣿⣧\033[0m ",
		"\033[32m\033[1m⠀⠀⠀⠀⠘⣿⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⣿⣷⣤⣀⣸⣧⣠⣤⣴⣶⣾⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⡿\033[0m ",
		"\033[32m\033[1m⠀⠀⠀⠀⠀⣿⣿⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⣿⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⣿⣿⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⣿⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⠇\033[0m ",
		"\033[32m\033[1m⠀⠀⠀⠀⠀⠘⢿⣿⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⣿⣿\033[2m⣿\033[0m\033[1m\033[32m⣿⣿⣿⣿⣿⠿⠟⠛⠉⠀\033[0m ",
		"\033[32m\033[1m⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠉⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠀\033[0m ",
	}

	moduleCount := len(listAllModules())
	prefix := "s"

	if moduleCount == 1 {
		prefix = ""
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

			if hovering == 6 { // Back
				clearTerminal()
				mainMenu(true)
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

			fmt.Printf("\033[%dA", 17)
			fmt.Print("\033[2K")

			// Print Menu
			fmt.Println("")
			fmt.Println(manowar[0] + " \033[0m\033[1m GhostShip v0.1.0\033[0m")
			fmt.Println(manowar[1] + "\033[0m by: \033]8;;https://reeceharris.net\033\\notreeceharris\033]8;;\033\\\033[32m\033[1m")
			fmt.Println(manowar[2])
			fmt.Println(manowar[3] + " " + " \033[2mA collection of " + fmt.Sprint(moduleCount) + " hacking tool" + prefix + ".\033[0m")
			fmt.Println(manowar[4])
			fmt.Println(manowar[5] + " " + colors[0] + cursor[0] + " all the gear \033[2m(all tools)" + "\033[0m")
			fmt.Println(manowar[6] + " " + colors[1] + cursor[1] + " open seas    \033[2m(open-source intelligence)" + "\033[0m")
			fmt.Println(manowar[7] + " " + colors[2] + cursor[2] + " plundering   \033[2m(exploitation)" + "\033[0m")
			fmt.Println(manowar[8] + " " + colors[3] + cursor[3] + " looting      \033[2m(post-explot)" + "\033[0m")
			fmt.Println(manowar[9] + " " + colors[4] + cursor[4] + " charting     \033[2m(reporting)" + "\033[0m")
			fmt.Println(manowar[10] + " " + colors[5] + cursor[5] + " scour        \033[2m(search)" + "\033[0m")
			fmt.Println(manowar[11] + " " + colors[6] + cursor[6] + " bout ship    \033[2m(go back)" + "\033[0m")
			fmt.Println(manowar[12])
			fmt.Println(manowar[13] + " \033[2m (k = up, j = down, space = submit) \033[0m")
			fmt.Println(manowar[14])
			fmt.Println(" ")
		}

		first_loop = false
	}
}

func mainMenu(init_clear bool) {

	var (
		hovering          = 0
		options           = []string{"module", "booty", "settings", "report", "exit"}
		colors            = make([]string, len(options))
		cursor            = make([]string, len(options))
		first_loop        = true
		clear             = init_clear
		b          []byte = make([]byte, 1)
	)

	ke := "\033[0m\033[32m\033[1m"
	kraken := []string{
		" ⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣴⣶⣤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀ ",
		" ⠀⠀⠀⠀⣠⡤⣤⣄⣾⣿⣿⣿⣿⣿⣿⣷⣠⣀⣄⡀⠀⠀⠀⠀ ",
		" ⠀⠀⠀⠀⠙⠀⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣬⡿⠀⠀⠀⠀ ",
		" ⠀⠀⠀⠀⠀⢀⣼⠟⢿⣿⣿⣿⣿⣿⣿⡿⠘⣷⣄⠀⠀⠀⠀⠀ ",
		" ⣰⠛⠛⣿⢠⣿⠋⠀⠀⢹⠻⣿⣿⡿⢻⠁⠀⠈⢿⣦⠀⠀⠀⠀ ",
		" ⢈⣵⡾⠋⣿⣯⠀⠀⢀⣼⣷⣿⣿⣶⣷⡀⠀⠀⢸⣿⣀⣀⠀⠀ ",
		" ⢾⣿⣀⠀⠘⠻⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⠿⣿⡁⠀⠀⠀ ",
		" ⠈⠙⠛⠿⠿⠿⢿⣿⡿⣿⣿⡿⢿⣿⣿⣿⣷⣄⠀⠘⢷⣆⠀⠀ ",
		" ⠀⠀⠀⠀⠀⢠⣿⠏⠀⣿⡏⠀⣼⣿⠛⢿⣿⣿⣆⠀⠀⣿⡇⡀ ",
		" ⠀⠀⠀⠀⢀⣾⡟⠀⠀⣿⣇⠀⢿⣿⡀⠈⣿⡌⠻⠷⠾⠿⣻⠁ ",
		" ⠀⠀⣠⣶⠟⠫⣤⠀⠀⢸⣿⠀⣸⣿⢇⡤⢼⣧⠀⠀⠀⢀⣿⠀ ",
		" ⠀⣾⡏⠀⡀⣠⡟⠀⠀⢀⣿⣾⠟⠁⣿⡄⠀⠻⣷⣤⣤⡾⠋⠀ ",
		" ⠀⠙⠷⠾⠁⠻⣧⣀⣤⣾⣿⠋⠀⠀⢸⣧⠀⠀⠀⠉⠁⠀⠀⠀ ",
		" ⠀⠀⠀⠀⠀⠀⠈⠉⠉⠹⣿⣄⠀⠀⣸⡿⠀⠀⠀⠀⠀⠀⠀⠀ ",
		" ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠛⠿⠟⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀ ",
	}

	// Put random spots to the kraken
	rand.Seed(time.Now().Unix())
	for i, val := range kraken {
		changed := 0
		tries := 0
		changedArr := []int{}
		chars := []rune(val)
		for changed < 1 && tries < 50 {
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
				err := openURL("https://github.com/ChatterSec/GhostShip/issues")
				if err != nil {
					clearTerminal()
					fmt.Printf("\033[%dA", 17)
					fmt.Print("\033[2K")
					fmt.Println("Issue opening issue url, please follow manually: https://github.com/ChatterSec/GhostShip/issues")
				}
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
			fmt.Println(ke + kraken[7] + "\033[0m" + colors[0] + cursor[0] + " man-o-war    \033[2m(tools)")
			fmt.Println(ke + kraken[8] + "\033[0m" + colors[1] + cursor[1] + " booty        \033[2m(0 captures)")
			fmt.Println(ke + kraken[9] + "\033[0m" + colors[2] + cursor[2] + " adjustments  \033[2m(settings)")
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

	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		exit()
	} else {

		if height < 19 {
			fmt.Println("Please resize your terminal to at least 19 lines. (currently", height, ")")
			exit()
		}

		if width < 85 {
			fmt.Println("Please resize your terminal to at least 85 columns. (currently", width, ")")
			exit()
		}
	}

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

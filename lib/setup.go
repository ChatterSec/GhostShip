package lib

import (
	"fmt"
	"os/exec"
)

func setup() {
	cmd := exec.Command("python", "--version")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Python is not installed")
	} else {
		fmt.Printf("Python is installed: %s", out)

		// Install Python requirements
		cmd = exec.Command("python", "-m", "pip", "install", "-r", "requirements.txt")
		err = cmd.Run()

		if err != nil {
			fmt.Println("Failed to install Python requirements")
		} else {
			fmt.Println("Python requirements installed successfully")

			// Change directory and run migrate
			cmd = exec.Command("python3", "manage.py", "migrate")
			cmd.Dir = "modules/sherlock"
			err = cmd.Run()

			if err != nil {
				fmt.Println("Failed to run migrate")
			} else {
				fmt.Println("Migrate ran successfully")
			}
		}
	}
}

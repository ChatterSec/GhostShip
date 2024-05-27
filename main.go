package main

import (
	"fmt"
	"ghost-ship/helper" // assuming the ascii.go file is in the same directory
)

func main() {
	asciiArt, err := helper.GetAsciiArt("key1") // assuming GetAsciiArt is a function in ascii.go that returns ascii art
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(asciiArt)
}

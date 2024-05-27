package helper

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var asciiArtDict = map[string]map[int]string{
	"key1": {
		80:  "ASCII Art1 for 80 width terminal",
		100: "ASCII Art1 for 100 width terminal",
	},
	"key2": {
		80:  "Another ASCII Art2 for 80 width terminal",
		100: "Another ASCII Art2 for 100 width terminal",
	},
	// Add more keys and terminal sizes as needed
}

func GetAsciiArt(key string) (string, error) {
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return "", err
	}

	// Check if the key exists in the dictionary
	artDict, ok := asciiArtDict[key]
	if !ok {
		return "", fmt.Errorf("no ASCII art found for key %s", key)
	}

	// Find the largest terminal size that is less than or equal to the current width
	maxSize := -1
	for size := range artDict {
		if size <= width && size > maxSize {
			maxSize = size
		}
	}

	// If no suitable size was found, return an error
	if maxSize == -1 {
		return "", fmt.Errorf("no ASCII art found for key %s and terminal width %d", key, width)
	}

	// Return the ASCII art for the found size
	return artDict[maxSize], nil
}

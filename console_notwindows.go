//go:build !windows
// +build !windows

package main

func setConsoleMode() {
	// This function does nothing on non-Windows platforms.
}

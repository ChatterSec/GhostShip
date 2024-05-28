//go:build windows
// +build windows

package main

import (
	"syscall"
	"unsafe"
)

const (
	ENABLE_ECHO_INPUT uint32 = 0x0004
	ENABLE_LINE_INPUT uint32 = 0x0002
)

func setConsoleMode() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode := kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode := kernel32.NewProc("SetConsoleMode")

	var mode uint32
	procGetConsoleMode.Call(uintptr(syscall.Stdin), uintptr(unsafe.Pointer(&mode)))
	mode &^= ENABLE_ECHO_INPUT
	mode &^= ENABLE_LINE_INPUT
	procSetConsoleMode.Call(uintptr(syscall.Stdin), uintptr(mode))
}

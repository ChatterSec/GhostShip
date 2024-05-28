//go:build windows
// +build windows

package main

import (
	"syscall"
	"unsafe"
)

const (
	STD_INPUT_HANDLE         = -10 & (1<<32 - 1)
	ENABLE_ECHO_INPUT uint32 = 0x0004
	ENABLE_LINE_INPUT uint32 = 0x0002
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
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

func resetConsoleMode() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode := kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode := kernel32.NewProc("SetConsoleMode")

	var mode uint32
	procGetConsoleMode.Call(uintptr(syscall.Stdin), uintptr(unsafe.Pointer(&mode)))
	mode |= ENABLE_ECHO_INPUT
	mode |= ENABLE_LINE_INPUT
	procSetConsoleMode.Call(uintptr(syscall.Stdin), uintptr(mode))
}

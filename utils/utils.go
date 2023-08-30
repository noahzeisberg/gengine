package utils

import (
  "syscall"
  "unsafe"
)

func SetTitle(title string) {
  handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return
	}
	syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
}

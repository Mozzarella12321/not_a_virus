package message

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

var (
	moduser32               = syscall.NewLazyDLL("user32.dll")
	procEnumWindows         = moduser32.NewProc("EnumWindows")
	procGetWindowText       = moduser32.NewProc("GetWindowTextW")
	procSetForegroundWindow = moduser32.NewProc("SetForegroundWindow")
)

func enumWindowsProc(hwnd syscall.Handle, lParam uintptr) uintptr {
	var title [256]uint16
	_, _, _ = procGetWindowText.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&title[0])), uintptr(len(title)))

	if strings.Contains(syscall.UTF16ToString(title[:]), *(*string)(unsafe.Pointer(lParam))) {
		*(*syscall.Handle)(unsafe.Pointer(lParam)) = hwnd
		return 0 // Stop enumeration
	}
	return 1 // Continue enumeration
}

func findWindowContainingTitle(partialTitle string) syscall.Handle {
	var hwnd syscall.Handle
	enumWindowsProcArg := uintptr(unsafe.Pointer(&hwnd))
	procEnumWindows.Call(syscall.NewCallback(enumWindowsProc), enumWindowsProcArg)
	return hwnd
}

func FocusWindow(partialTitle string) error {
	hwnd := findWindowContainingTitle(partialTitle)

	if hwnd != 0 {
		// Фокусируем окно
		_, _, err := procSetForegroundWindow.Call(uintptr(hwnd))
		if err != nil {
			return fmt.Errorf("не удалось установить фокус на окно Notepad: %v", err)
		}
	}
	return nil
}

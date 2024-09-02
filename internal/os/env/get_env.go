package env

import (
	"fmt"
	"log"
	"os/user"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

var CurrentOS string
var CurrentUser *user.User
var CurrentName string

// CurrentEnv string  //
// CurrentUser *user.User  //
// CurrentName string
func GetEnvironmentData() {
	getOS()
	getCurrentuser()
	assignAndFormatUsername()
}

// assigns system name to global var CurrentEnv
func getOS() {
	CurrentOS = runtime.GOOS
}

func getCurrentuser() {
	var err error
	CurrentUser, err = user.Current()
	if err != nil {
		log.Fatalf("Ошибка при получении пользователя: %v", err)
	}
}

var defaultUsernames = map[string]bool{
	"administrator": true,
	"user":          true,
	"defaultuser":   true,
	"guest":         true,
	"admin":         true,
	"owner":         true,
}

func assignAndFormatUsername() {
	var err error
	CurrentName, err = getCurrentUserName()
	if err != nil {
		log.Println(err)
	}
	if defaultUsernames[strings.ToLower(CurrentName)] {
		uname := strings.Title(strings.ToLower(CurrentUser.Username))
		if strings.Contains(uname, "\\") {
			parts := strings.Split(uname, "\\")
			CurrentName = parts[0]
		} else {
			CurrentName = uname
		}
	}

}

// func getCurrentUserName1() (string, error) {
// 	var username [256]uint16
// 	var size uint32 = uint32(len(username))

// 	// Получаем имя текущего пользователя
// 	err := syscall.GetUserNameEx(&username[0], &size)
// 	if err != nil {
// 		return "", fmt.Errorf("не удалось получить имя пользователя: %v", err)
// 	}

// 	// Преобразуем имя пользователя из UTF-16 в строку
// 	return syscall.UTF16ToString(username[:]), nil
// }

func getCurrentUserName() (string, error) {
	var username [256]uint16
	var size uint32 = uint32(len(username))

	// Загрузка библиотеки advapi32.dll и получение адреса функции GetUserNameW
	modadvapi32 := syscall.NewLazyDLL("advapi32.dll")
	procGetUserName := modadvapi32.NewProc("GetUserNameW")

	// Вызов функции GetUserNameW
	ret, _, err := procGetUserName.Call(uintptr(unsafe.Pointer(&username[0])), uintptr(unsafe.Pointer(&size)))
	if ret == 0 {
		return "", fmt.Errorf("не удалось получить имя пользователя: %v", err)
	}

	// Преобразуем имя пользователя из UTF-16 в строку
	return syscall.UTF16ToString(username[:]), nil
}

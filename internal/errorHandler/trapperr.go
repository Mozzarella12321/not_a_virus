package e

import (
	"fmt"
	"log"
	"os"
	"time"
)

var CheckLogs = false

func HandleF(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Handle(e error) {
	if e == nil {
		return
	}
	logError(e)
	log.Println(e)
}

func logError(err error) {
	if err == nil {
		return
	}
	CheckLogs = true
	// Открываем файл Logs.txt для добавления информации
	file, fileErr := os.OpenFile("Logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if fileErr != nil {
		fmt.Printf("Не удалось открыть файл журнала: %v\n", fileErr)
		return
	}
	defer file.Close()

	// Форматируем ошибку с меткой времени
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("%s - ERROR: %v\n", timestamp, err)

	// Записываем ошибку в файл
	_, writeErr := file.WriteString(logMessage)
	if writeErr != nil {
		fmt.Printf("Не удалось записать в файл журнала: %v\n", writeErr)
	}
}

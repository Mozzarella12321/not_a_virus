package gc

import (
	"bufio"
	"fmt"
	"os"
)

// appendToFile создаёт файл, если его нет, и добавляет данные с новой строки
func AppendToFile(content string) error {
	// Открываем файл в режиме добавления (append) или создаём его, если он не существует
	file, err := os.OpenFile("tempFiles.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем содержимое в файл с новой строки
	if _, err := file.WriteString(content + "\n"); err != nil {
		return err
	}

	return nil
}

func HandleRemoval(content string) error {
	return AppendToFile(content)
}

func DeleteFilesFromList(fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
		os.Remove(fileName)
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileToDelete := scanner.Text()
		if err := os.Remove(fileToDelete); err != nil {
			fmt.Printf("Ошибка при удалении файла %s: %v\n", fileToDelete, err)
		} else {
			fmt.Printf("Файл %s успешно удален.\n", fileToDelete)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

package message

import (
	"log"
	"os"
	"os/exec"
	e "virus/internal/errorHandler"
	"virus/internal/os/env"
	"virus/internal/os/gc"

	"github.com/pkg/errors"
)

func WriteTxtToUser(s string) error {
	switch env.CurrentOS {
	case "windows":
		// Создаем временный файл с текстом
		tmpFile, err := os.CreateTemp("", "Hello, "+env.CurrentName+"     12321*")
		if err != nil {
			return errors.Errorf("Ошибка создания временного файла: %v", err)
		}
		defer e.Handle(gc.HandleRemoval(tmpFile.Name())) // Удаляем файл после завершения программы

		//text := []byte(fmt.Sprintf("Привет, %s", env.CurrentName))
		// Пишем текст в файл
		if _, err := tmpFile.Write([]byte(s)); err != nil {
			return errors.Errorf("Ошибка записи в файл: %v", err)
		}

		// Закрываем файл, чтобы Notepad смог его открыть
		if err := tmpFile.Close(); err != nil {
			return errors.Errorf("Ошибка закрытия файла: %v", err)
		}

		// Открываем Notepad с нашим файлом
		cmd := exec.Command("notepad", tmpFile.Name())
		//cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Start-Process notepad -ArgumentList '%s'", tmpFile.Name()))
		err = cmd.Start()
		if err != nil {
			return errors.Errorf("Ошибка запуска Notepad: %v", err)
		}
		FocusWindow(tmpFile.Name())
		//FocusWindow(tmpFile.Name())
		//Ждем завершения Notepad (необязательно)
		// err = cmd.Wait()
		// if err != nil {
		// 	return errors.Errorf("Notepad завершился с ошибкой: %v", err)
		// }
	default:
		log.Fatalf("ты на чем это запускаешь ебик D;")
		//пока не придумал че с другими осями делать может и ну нахуй
	}
	return nil
}

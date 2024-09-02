package main

import (
	"fmt"
	"time"
	e "virus/internal/errorHandler"
	"virus/internal/os/env"
	"virus/internal/os/gc"
	"virus/internal/os/message"
)

func main() {
	fmt.Println("Не закрывай пж, я сам закроюсь попозже")
	env.GetEnvironmentData()
	//Если нет тг открывается окно, "выбрать приложение", а не err
	//и по какой-то причине у некоторых просто не открывается тг
	err := message.OpenURL("tg://resolve?domain=mozzarella12321blya")
	if err != nil {
		e.Handle(err)
		e.Handle(message.OpenURL("https://t.me/mozzarella12321blya"))
	}
	time.Sleep(2000 * time.Millisecond)
	e.Handle(message.WriteTxtToUser("Подпишись, а иначе..."))
	time.Sleep(5000 * time.Millisecond)
	e.Handle(message.WriteTxtToUser("А иначе хуй я че сделаю иначе \nу меня скила не хватает"))
	time.Sleep(1000 * time.Millisecond)
	if e.CheckLogs {
		e.Handle(message.WriteTxtToUser("Походу залупа какая-то, скинь мне логи пжпж"))
	}
	time.Sleep(2 * time.Second)
	defer gc.DeleteFilesFromList("tempFiles.txt")
}

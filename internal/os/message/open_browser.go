package message

import (
	"log"
	"os/exec"
	"virus/internal/os/env"
)

func OpenURL(url string) error {
	switch env.CurrentOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "start", url)
		return cmd.Start()
	default:
		log.Fatalf("ты на чем это запускаешь ебик D;")
	}
	return nil
}

package services

import (
	"os/exec"
)

func GenerateBaseApp(name string) error {
	cmd := exec.Command("flutter", "create", name)
	err := cmd.Run()
	return err
}

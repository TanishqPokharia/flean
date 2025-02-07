package services

import (
	"errors"
	"flean/constants"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Pubspec struct {
	Dependencies map[string]any `yaml:"dependencies"`
}

func AddDependencies(projectName string, deps ...string) error {
	if len(deps) == 0 {
		return nil
	}
	file, err := os.ReadFile(fmt.Sprintf("./%s/pubspec.yaml", projectName))
	if err != nil {
		return err
	}
	content := string(file)
	newDependencies := "dependencies:\n  "
	for _, v := range deps {
		newDependencies += fmt.Sprintf("%s:\n  ", v)
	}
	content = strings.Replace(content, "dependencies:", newDependencies, 1)
	err = os.WriteFile(fmt.Sprintf("./%s/pubspec.yaml", projectName), []byte(content), os.FileMode(constants.FilePermission))
	if err != nil {
		return err
	}
	return updateDependencies(projectName)
}

func updateDependencies(projectName string) error {
	err := os.Chdir(fmt.Sprintf("./%s", projectName))
	if err != nil {
		return errors.New("failed to updated dependencies, try running flutter pub get manually later\"")
	}
	cmd := exec.Command("flutter", "pub", "get")
	if err = cmd.Run(); err != nil {
		return errors.New("failed to updated dependencies, try running flutter pub get manually later")
	}
	if err = os.Chdir("../"); err != nil {
		return err
	}
	return err
}

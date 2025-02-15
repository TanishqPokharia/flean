package services

import (
	"errors"
	"flean/constants"
	"flean/tui"
	"flean/utils"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
)

func GenerateClean(name string) {
	generateData(fmt.Sprintf("./%s/lib", name))
	generateDomain(fmt.Sprintf("./%s/lib", name))
	generatePresentation(fmt.Sprintf("./%s/lib", name))
	generateThemes(name)
	generateUtils(name)
}

func GenerateFFBase(name string) {
	generateThemes(name)
	generateUtils(name)
	generateFeatures(name)
}

func GenerateFFClean(path string, name string) {
	_, err := os.Stat(path)
	dir, err := os.Stat(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		utils.LogFatalError(err)
	}
	if dir != nil && dir.IsDir() {
		fmt.Printf(lipgloss.NewStyle().Foreground(tui.Red).Render(fmt.Sprintf("%s feature already exists\n", name)))
		return
	}
	generateData(path)
	generateDomain(path)
	generatePresentation(path)
	fmt.Println(tui.LogStyle.Render(fmt.Sprintf("%s feature added\n", name)))
}

func RemoveFeatureClean(name string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	details, err := GetProjectDetails(pwd)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/lib/features/%s", details.Directory, details.Name, name)
	_, err = os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal(lipgloss.NewStyle().Foreground(tui.Red).Render(fmt.Sprintf("%s feature does not exist", name)))
		} else {
			utils.LogFatalError(err)
		}
	}
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func generateData(path string) {
	err := os.MkdirAll(fmt.Sprintf("%s/data", path), os.FileMode(constants.FilePermission))
	if err != nil {
		utils.LogFatalError(err)
	}
	utils.GenerateNested(fmt.Sprintf("%s/data/datasource", path))
	utils.GenerateNested(fmt.Sprintf("%s/data/repository", path))
	utils.GenerateNested(fmt.Sprintf("%s/data/models", path))
}

func generateDomain(path string) {
	err := os.MkdirAll(fmt.Sprintf("%s/domain", path), os.FileMode(constants.FilePermission))
	if err != nil {
		utils.LogFatalError(err)
	}
	utils.GenerateNested(fmt.Sprintf("%s/domain/entities", path))
	utils.GenerateNested(fmt.Sprintf("%s/domain/usecases", path))
	utils.GenerateNested(fmt.Sprintf("%s/domain/repository", path))
}

func generatePresentation(path string) {
	err := os.MkdirAll(fmt.Sprintf("%s/presentation", path), os.FileMode(constants.FilePermission))
	if err != nil {
		utils.LogFatalError(err)
	}
	utils.GenerateNested(fmt.Sprintf("%s/presentation/widgets", path))
	utils.GenerateNested(fmt.Sprintf("%s/presentation/screens", path))
	utils.GenerateNested(fmt.Sprintf("%s/presentation/bloc", path))
	utils.GenerateNested(fmt.Sprintf("%s/presentation/cubit", path))
}

func generateThemes(path string) {
	utils.GenerateNested(fmt.Sprintf("./%s/lib/themes", path))
	_, err := os.Create(fmt.Sprintf("./%s/lib/themes/colors.dart", path))
	if err != nil {
		utils.LogFatalError(err)
	}
	_, err = os.Create(fmt.Sprintf("./%s/lib/themes/icons.dart", path))
	if err != nil {
		utils.LogFatalError(err)
	}
	_, err = os.Create(fmt.Sprintf("./%s/lib/themes/themes.dart", path))
	if err != nil {
		utils.LogFatalError(err)
	}
}

func generateUtils(path string) {
	err := os.Mkdir(fmt.Sprintf("./%s/lib/utils", path), os.FileMode(constants.FilePermission))
	if err != nil {
		utils.LogFatalError(err)
	}
}

func generateFeatures(path string) {
	err := os.MkdirAll(fmt.Sprintf("./%s/lib/features", path), os.FileMode(constants.FilePermission))
	if err != nil {
		utils.LogFatalError(err)
	}
}

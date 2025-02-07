package services

import (
	"errors"
	"flean/tui"
	"flean/utils"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
	"strings"
)

func RemoveBloc(name string) {
	pwd, err := os.Getwd()
	style := lipgloss.NewStyle()
	if err != nil {
		log.Fatal(err)
	}
	details, err := GetProjectDetails(pwd)
	if err != nil {
		log.Fatal(err)
	}
	rootPath := details.Directory + "/" + details.Name
	lower := strings.ToLower(name)
	path := fmt.Sprintf("%s/lib%s/%s", rootPath, details.BlocPath, lower)
	_, err = os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal(style.Foreground(tui.Red).Render(fmt.Sprintf("%s bloc does not exist", name)))
		} else {
			utils.LogFatalError(err)
		}
	}
	err = os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s bloc removed\n", name)
}

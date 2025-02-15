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

func RemoveBloc(name string, feature string) {
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

	// handle path for the bloc in a specific feature
	lower := strings.ToLower(name)

	// remove the bloc
	path := ""
	if feature != "" {
		path = fmt.Sprintf("%s/lib/features/%s/%s/%s", rootPath, feature, details.BlocPath, lower)
	} else {
		path = fmt.Sprintf("%s/lib%s/%s", rootPath, details.BlocPath, lower)
	}
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
	fmt.Print(tui.LogStyle.Render(fmt.Sprintf("%s bloc removed\n", name)))
}

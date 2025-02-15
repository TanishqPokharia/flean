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

func RemoveCubit(name string, feature string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	details, err := GetProjectDetails(pwd)
	if err != nil {
		log.Fatal(err)
	}
	rootPath := details.Directory + "/" + details.Name
	// handle path for the bloc in a specific feature
	lower := ""
	if feature != "" {
		lower += fmt.Sprintf("/features/%s/", strings.ToLower(feature))
	}
	lower += strings.ToLower(name)

	// add the bloc to constructed path
	cubitPath := strings.Split(details.BlocPath, "/bloc")[0] + "/cubit"
	path := fmt.Sprintf("%s/lib%s/%s", rootPath, cubitPath, lower)
	_, err = os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal(lipgloss.NewStyle().Foreground(tui.Red).Render(fmt.Sprintf("%s cubit does not exist", name)))
		} else {
			utils.LogFatalError(err)
		}
	}
	err = os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(tui.LogStyle.Render(fmt.Sprintf("%s cubit removed\n", name)))
}

package services

import (
	"flean/constants"
	"flean/templates/cubit"
	"flean/tui"
	"flean/utils"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func GenerateCubit(name string, feature string) {
	pwd, err := os.Getwd()
	if err != nil {
		utils.LogFatalError(err)
	}
	details, err := GetProjectDetails(pwd)
	if err != nil {
		utils.LogFatalError(err)
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
	if os.IsExist(err) {
		log.Fatalf("%s cubit already exists", name)
	}
	err = os.MkdirAll(path, os.FileMode(constants.FilePermission))
	if err != nil {
		utils.LogFatalError(err)
	}
	generateState(name, path, cubit.CubitStateTemplate)
	generateCubit(name, path, cubit.CubitTemplate)
	fmt.Print(tui.LogStyle.Render(fmt.Sprintf("%s cubit added\n", name)))
}

func generateCubit(name string, path string, tmpl string) {
	tmp, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		utils.LogFatalError(err)
	}
	lower := strings.ToLower(name)
	file, err := os.Create(fmt.Sprintf("%s/%s_cubit.dart", path, lower))
	defer func() {
		err := file.Close()
		if err != nil {
			utils.LogFatalError(err)
		}
	}()
	if err != nil {
		utils.LogFatalError(err)
	}
	err = tmp.Execute(file, utils.TemplateData{
		Upper: strings.ToUpper(name[:1]) + name[1:],
		Lower: strings.ToLower(name),
	})
	if err != nil {
		utils.LogFatalError(err)
	}
}

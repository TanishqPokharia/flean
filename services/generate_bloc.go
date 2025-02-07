package services

import (
	"errors"
	"flean/constants"
	"flean/templates/bloc"
	"flean/tui"
	"flean/utils"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func GenerateBloc(name string) {
	pwd, err := os.Getwd()
	if err != nil {
		utils.LogFatalError(err)
	}
	details, err := GetProjectDetails(pwd)
	if err != nil {
		utils.LogFatalError(err)
	}
	rootPath := details.Directory + "/" + details.Name
	lower := strings.ToLower(name)
	path := fmt.Sprintf("%s/lib%s/%s", rootPath, details.BlocPath, lower)
	dir, err := os.Stat(path)
	if !errors.Is(err, os.ErrNotExist) {
		utils.LogFatalError(err)
	}
	if dir != nil && dir.IsDir() {
		fmt.Printf("%s bloc already exists\n", name)
		return
	}
	err = os.MkdirAll(path, os.FileMode(constants.FilePermission))
	if err != nil {
		utils.LogFatalError(err)
	}
	generateState(name, path, bloc.BlocStateTemplate)
	generateEvent(name, path, bloc.BlocEventTemplate)
	generateBloc(name, path, bloc.BlocTemplate)
	fmt.Print(tui.LogStyle.Render(fmt.Sprintf("%s bloc added\n", name)))
}

func generateState(name string, path string, tmpl string) {
	tmp, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		utils.LogFatalError(err)
	}
	lower := strings.ToLower(name)
	file, err := os.Create(fmt.Sprintf("%s/%s_state.dart", path, lower))
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

func generateEvent(name string, path string, tmpl string) {
	tmp, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		utils.LogFatalError(err)
	}
	lower := strings.ToLower(name)
	file, err := os.Create(fmt.Sprintf("%s/%s_event.dart", path, lower))
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

func generateBloc(name string, path string, tmpl string) {
	tmp, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		utils.LogFatalError(err)
	}
	lower := strings.ToLower(name)
	file, err := os.Create(fmt.Sprintf("%s/%s_bloc.dart", path, lower))
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

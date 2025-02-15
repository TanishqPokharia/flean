package services

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type FleanYaml struct {
	Name         string `yaml:"name"`
	Directory    string `yaml:"directory"`
	Architecture string `yaml:"architecture"`
	BlocPath     string `yaml:"bloc_path"`
}

func StoreProjectDetails(projectName string, arch string) error {
	file, err := os.Create(fmt.Sprintf("./%s/flean.yaml", projectName))
	if err != nil {
		return err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return errors.New("warning! Could not get directory details, please add the directory path manually")
	}
	config := FleanYaml{Name: projectName, Directory: pwd, Architecture: arch}
	if arch == "Clean" || arch == "Feature First Clean" {
		config.BlocPath = "/presentation/bloc"
	}
	if arch == "MVVM" || arch == "MVC" {
		config.BlocPath = "/bloc"
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	if _, err = file.Write(data); err != nil {
		return err
	}
	return nil
}

func GetProjectDetails(pwd string) (FleanYaml, error) {
	pathTillLib := strings.Split(pwd, "lib")[0]
	file, err := os.ReadFile(fmt.Sprintf("%s/flean.yaml", pathTillLib))
	if err != nil {
		return FleanYaml{}, err
	}
	var projectDetails FleanYaml
	if err = yaml.Unmarshal(file, &projectDetails); err != nil {
		return FleanYaml{}, err
	}
	return projectDetails, nil
}

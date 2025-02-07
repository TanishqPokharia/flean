package services

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func RemoveCubit(name string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	details, err := GetProjectDetails(pwd)
	if err != nil {
		log.Fatal(err)
	}
	rootPath := details.Directory + "/" + details.Name
	lower := strings.ToLower(name)
	cubitPath := strings.Split(details.BlocPath, "/bloc")[0] + "/cubit"
	path := fmt.Sprintf("%s/lib%s/%s", rootPath, cubitPath, lower)
	err = os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s cubit removed\n", name)
}

package services

import (
	"flean/constants"
	"flean/utils"
	"fmt"
	"os"
)

func GenerateMVC(name string) {
	folders := []string{"models", "controllers", "views", "services"}
	for _, v := range folders {
		err := os.MkdirAll(fmt.Sprintf("./%s/lib/%s", name, v), os.FileMode(constants.FilePermission))
		if err != nil {
			utils.LogFatalError(err)
		}
	}
}

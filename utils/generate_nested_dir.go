package utils

import (
	"flean/constants"
	"log"
	"os"
)

func GenerateNested(path string) {
	err := os.MkdirAll(path, os.FileMode(constants.FilePermission))
	if err != nil {
		log.Fatal(err)
	}
}

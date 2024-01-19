package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ncruces/zenity"
)

func main(){
	inDir, err := zenity.SelectFile(
		zenity.Filename(""),
		zenity.Directory(),
		zenity.DisallowEmpty(),
		zenity.Title("Select input directory"),
	)
	if err != nil {
		zenity.Error(
			err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon,
		)
		log.Fatal(err)
	}

	dlm, err := zenity.Entry("Enter delimitier",
		zenity.Title("Delimiter"),
	)


	if dlm != "" {
		// loop over folders
		files, err := os.ReadDir(inDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if file.IsDir() {
				newName := strings.TrimSpace(file.Name()[strings.Index(file.Name(), dlm)+1:])+filepath.Ext(file.Name())
				fmt.Println(file.Name(), newName)
			}
		}

		zenity.Info("Folders renamed!",
			zenity.Title("Complete"),
			zenity.InfoIcon,
		)
	} else {
		zenity.Info("No delimiter provided!",
			zenity.Title("Attention"),
			zenity.InfoIcon,
		)
	}
}
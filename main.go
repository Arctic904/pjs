package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/arctic904/pjs/project"
	"github.com/arctic904/pjs/tui"
	"github.com/arctic904/pjs/tui/constants"
	"github.com/arctic904/pjs/utils"
	"github.com/pkg/errors"
)

func getLocalFolder() string {
	dirPath := fmt.Sprintf("%v/.pjs", xdg.DataHome)
	_, err := os.ReadDir(dirPath)
	if err != nil {
		err = os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = os.WriteFile(fmt.Sprintf("%v/proj.json", dirPath), []byte("[]"), 0777)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error writing project file"))
	}
	return dirPath
}

func main() {
	path := getLocalFolder()
	constants.Path = path
	projects, err := utils.ReadProjJson()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error reading project file"))
	}

	if len(projects) < 1 {
		name := project.NewProjectPrompt()
		_, err := project.CreateProject(name, path)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error creating project"))
		}
	} else {
		tui.StartTea(projects)
	}
}

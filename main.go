package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
)

func getLocalFolder() string {
	dirPath := fmt.Sprintf("%v/.pjs", xdg.DataHome)
	_, err := os.ReadDir(dirPath)
	if err != nil {
		ferr := os.Mkdir(dirPath, os.ModePerm)
		if ferr != nil {
			log.Fatal(ferr)
		}
	}
	return dirPath
}

func main() {
	path := getLocalFolder()
	projects, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	// if len(projects) < 1 {
	// 	name := project.NewProjectPrompt()
	// 	_, err := pr.CreateProject(name)
	// 	if err != nil {
	// 		log.Fatal(errors.Wrap(err, "error creating project"))
	// 	}
	// } else {
	// 	tui.StartTea(pr, er)
	// }
	// db := openSqlite()
	// pr := project.GormRepository{DB: db}
	// er := entry.GormRepository{DB: db}
	// projects, err := pr.GetAllProjects()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if len(projects) < 1 {
	// 	name := project.NewProjectPrompt()
	// 	_, err := pr.CreateProject(name)
	// 	if err != nil {
	// 		log.Fatal(errors.Wrap(err, "error creating project"))
	// 	}
	// } else {
	// 	tui.StartTea(pr, er)
	// }
}

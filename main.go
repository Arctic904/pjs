package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/arctic904/pjs/project"
	// "github.com/arctic904/pjs/tui"
	// "github.com/pkg/errors"
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
	println(projects)
	file, err := os.ReadFile("./example_config.json")
	if err != nil {
		log.Fatal(err)
	}
	var projectList []project.Project
	MyJson := []byte(file)
	err = json.Unmarshal(MyJson, &projectList)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", projectList)
	println("")
	// if len(projects) < 1 {
	// 	name := project.NewProjectPrompt()
	// 	_, err := project.CreateProject(name)
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

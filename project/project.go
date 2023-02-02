package project

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/arctic904/pjs/utils"
)

const (
	format string = "%v\n"
)

// Project the project holds entries
type Project struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

// NewProject create a new project instance.
// DeletedAt defaults to the zero value for time.Time.
func NewProject(id uint, name string) *Project {
	return &Project{Name: name}
}

// Implement list.Item for Bubbletea TUI

// Title the project title to display in a list
func (p Project) Title() string { return p.Name }

// Description the project description to display in a list
func (p Project) Description() string { return fmt.Sprintf("%v", p.Desc) }

// FilterValue choose what field to use for filtering in a Bubbletea list component
func (p Project) FilterValue() string { return p.Name }

// Repository CRUD operations for Projects
type Repository interface {
	PrintProjects()
	HasProjects() bool
	GetProjectByID(projectID uint) (Project, error)
	GetAllProjects() ([]Project, error)
	CreateProject(name string) (Project, error)
	DeleteProject(projectID uint) error
	RenameProject(projectID uint) error
}

// GetProjectByID get a project by ID
func GetProjectByID(projectID int, path string) (utils.Project, error) {
	var preturn utils.Project
	projects, err := utils.ReadProjJson()
	if err != nil {
		return preturn, fmt.Errorf("Cannot find project: %v", err)
	}

	return projects[projectID], nil
}

// PrintProjects print all projects to the console
func PrintProjects(path string) {
	projects, err := GetAllProjects(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects {
		fmt.Printf(format, project.Name)
	}
}

// GetAllProjects retrieve all projects from the database
func GetAllProjects(path string) ([]utils.Project, error) {
	projects, err := utils.ReadProjJson()
	if err != nil {
		return projects, fmt.Errorf("Table is empty: %v", err)
	}
	return projects, nil
}

// HasProjects see if a database has any projects
func HasProjects(path string) bool {
	if projects, _ := utils.ReadProjJson(); len(projects) == 0 {
		return false
	}
	return true
}

// CreateProject add a new project to the database
func CreateProject(name string, path string) (Project, error) {
	proj := Project{Name: name}
	err := os.Mkdir(fmt.Sprintf("%v/%v", path, name), 0644)
	if err != nil {
		return proj, fmt.Errorf("Cannot create project: %v", err)
	}
	return proj, nil
}

// DeleteProject delete a project by ID
func DeleteProject(projectID int, path string) error {
	proj, err := GetAllProjects(path)
	if err != nil {
		return fmt.Errorf("Cannot delete project: %v", err)
	}
	name := proj[projectID].Name
	err = os.RemoveAll(fmt.Sprintf("%v/%v", path, name))
	remove(proj, projectID)
	err = utils.WriteProjJson([]byte(fmt.Sprintf("%v", proj)), path)
	if err != nil {
		return fmt.Errorf("Cannot delete project: %v", err)
	}
	return nil
}

// RenameProject rename an existing project
func RenameProject(id uint, name string, path string) {
	proj, err := GetAllProjects(path)
	if err != nil {
		log.Default().Printf("Cannot rename project: %v\n", err)
	}
	proj[id].Name = name
	err = utils.WriteProjJson([]byte(fmt.Sprintf("%v", proj)), path)
	if err != nil {
		log.Default().Printf("Cannot rename project: %v\n", err)
	}
}

// NewProjectPrompt create a new project from user input to console
func NewProjectPrompt() string {
	var name string
	fmt.Println("what would you like to name your project?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	return name
}

func remove(slice []utils.Project, s int) []utils.Project {
	return append(slice[:s], slice[s+1:]...)
}

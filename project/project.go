package project

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	format string = "%d : %s\n"
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
func (p Project) Description() string { return fmt.Sprintf("%d", p.ID) }

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
func GetProjectByID(projectID uint, path string) (Project, error) {
	var preturn Project
	files, err := os.ReadDir(fmt.Sprintf("%v/%v", path, projectID))
	if err != nil {
		return preturn, fmt.Errorf("Cannot find project: %v", err)
	}

	return preturn, nil
}

// PrintProjects print all projects to the console
func PrintProjects(path string) {
	projects, err := g.GetAllProjects()
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects {
		fmt.Printf(format, project.ID, project.Name)
	}
}

// GetAllProjects retrieve all projects from the database
func GetAllProjects(path string) ([]Project, error) {
	var projects []Project
	if err := g.DB.Find(&projects).Error; err != nil {
		return projects, fmt.Errorf("Table is empty: %v", err)
	}
	return projects, nil
}

// HasProjects see if a database has any projects
func HasProjects(path string) bool {
	if projects, _ := utils.readProjJson(); len(projects) == 0 {
		return false
	}
	return true
}

// CreateProject add a new project to the database
func CreateProject(name string, path string) (Project, error) {
	proj := Project{Name: name}
	if err := g.DB.Create(&proj).Error; err != nil {
		return proj, fmt.Errorf("Cannot create project: %v", err)
	}
	return proj, nil
}

// DeleteProject delete a project by ID
func DeleteProject(projectID uint, path string) error {
	if err := g.DB.Delete(&Project{}, projectID).Error; err != nil {
		return fmt.Errorf("Cannot delete project: %v", err)
	}
	return nil
}

// RenameProject rename an existing project
func RenameProject(id uint, name string, path string) {
	var newProject Project
	if err := g.DB.Where("id = ?", id).First(&newProject).Error; err != nil {
		log.Fatalf("Unable to rename project: %q", err)
	}
	newProject.Name = name
	if err := g.DB.Save(&newProject).Error; err != nil {
		log.Fatalf("Unable to save project: %q", err)
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

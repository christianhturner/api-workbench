package project

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/christianhturner/api-workbench/server"
	"gorm.io/gorm"
)

// The Project will reprepesnt a Host, and may expand this to support proxies
// in the future. Currently, a new project will represent two things in our
// project. Currently this will provide a human readable name for the project,
// the wwwPath, which is the path to path representation of the API Path.
// It will also represent the configPath, and this will feature a directory
// with the Project.name as the name and within will feature the API configuration
// available in this project. The API configurations will map to the physical
// file drectory representations.
type Project struct {
	gorm.Model
	Name        string
	wwwPath     string
	projectPath string
}

// Implement list.Item methods for bubbletea TUI

// List project name and implements list.Item method
func (p Project) Title() string { return p.Name }

// List project description for our list
func (p Project) Description() string { return fmt.Sprintf("%d", p.ID) }

// FilterValue defines what field is used for filtering and required method for list.Item
func (p Project) FilterValue() string { return p.Name }

type DBUtils interface {
	PrintProjects()
	CreateDBEntry(name string) (Project, error)
	HasProjects() bool
	GetAllProjects() ([]Project, error)
	GetProjectByID(projectID uint) (Project, error)
	DeleteProject(projectID uint) error
	RenameProject(projectID uint) error
}

type DB struct {
	DB *gorm.DB
}

func (d *DB) CreateDBEntry(name string) (Project, error) {
	dataDir, err := server.NewDataDir()
	if err != nil {
		log.Panic(err)
	}
	project := Project{
		Name:        name,
		wwwPath:     filepath.Join(dataDir.GetWWWPath(), name),
		projectPath: filepath.Join(dataDir.GetProjectPath(), name),
	}
	if err = d.DB.Create(&project).Error; err != nil {
		log.Printf("cannont create project: %v", err)
		return project, fmt.Errorf("cannont create project: %v", err)
	}
	project.createFilePathDependencies()
	return project, nil
}

func (p *Project) createFilePathDependencies() error {
	// Create data directory project Path for API configs
	_, err := os.Stat(p.projectPath)
	if os.IsNotExist(err) {
		log.Printf("Creating www sub directory for %s at %s", p.Name, p.projectPath)
		err = os.Mkdir(p.projectPath, 0755)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Path, %s already exist for project %s.", p.projectPath, p.Name)
	}
	// Create data directory www Path for emulated http server
	_, err = os.Stat(p.wwwPath)
	if os.IsNotExist(err) {
		log.Printf("Creating www sub-directory for %s at %s.", p.Name, p.wwwPath)
		err = os.Mkdir(p.wwwPath, 0755)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Path, %s already exist for project %s.", p.wwwPath, p.Name)
	}
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// GetProjectByID get project by ID
func (d *DB) GetProjectByID(projectID uint) (Project, error) {
	var project Project
	if err := d.DB.Where("id = ?", projectID).First(&project).Error; err != nil {
		log.Printf("cannont find project: %v", err)
		return project, fmt.Errorf("cannot find project: %v", err)
	}
	return project, nil
}

// GetAllPRojects retrieve all projects from database
func (d *DB) GetAllProjects() ([]Project, error) {
	var projects []Project
	if err := d.DB.Find(&projects).Error; err != nil {
		log.Printf("cannot find project: %v", err)
		return projects, fmt.Errorf("cannot find project: %v", err)
	}
	return projects, nil
}

// PrintProjects print all projects to the console
func (d *DB) PrintProjects() {
	projects, err := d.GetAllProjects()
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects {
		fmt.Printf("%d : %s\n", project.ID, project.Name)
	}
}

// HasProjects see if a database has any projects
func (d *DB) HasProjects() bool {
	if projects, _ := d.GetAllProjects(); len(projects) == 0 {
		return false
	}
	return true
}

// DeleteProject delete project by ID
func (d *DB) DeleteProject(projectID uint) error {
	if err := d.DB.Delete(&Project{}, projectID).Error; err != nil {
		log.Printf("cannot delete project: %v", err)
		return fmt.Errorf("cannot delete project: %v", err)
	}
	return nil
}

// RenameProject rename an existing project
func (d *DB) RenameProject(id uint, name string) error {
	var newProject Project
	if err := d.DB.Where("id = ?", id).First(&newProject).Error; err != nil {
		log.Printf("unable to rename project: %v", err)
		return fmt.Errorf("unable to rename project: %v", err)
	}
	newProject.Name = name
	if err := d.DB.Save(&newProject).Error; err != nil {
		log.Printf("unable to rename project: %v", err)
		return fmt.Errorf("unable to rename project: %v", err)
	}
	return nil
}

func NewProjectPrompt() string {
	var name string
	fmt.Println("Please provide a name for your project.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	return name
}

// \\\\ DEPRECIATED //// \\

func New(name string) (*Project, error) {
	dataDir, err := server.NewDataDir()
	if err != nil {
		log.Panic(err)
	}
	projectWwwPath := filepath.Join(dataDir.GetWWWPath(), name)
	_, err = os.Stat(projectWwwPath)
	if os.IsNotExist(err) {
		log.Printf("Creating www sub directory for %s at %s", name, projectWwwPath)
		err = os.Mkdir(projectWwwPath, 0755)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Path, %s already exist for project %s.", projectWwwPath, name)
	}
	projectProjectPath := filepath.Join(dataDir.GetProjectPath(), name)
	_, err = os.Stat(projectProjectPath)
	if os.IsNotExist(err) {
		log.Printf("Creating config sub directory for %s at %s", name, projectProjectPath)
		err = os.Mkdir(projectProjectPath, 0755)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Path, %s already exist for project %s.", projectProjectPath, name)
	}
	return &Project{
		Name:        name,
		wwwPath:     dataDir.GetWWWPath(),
		projectPath: dataDir.GetProjectPath(),
	}, nil
}

// func GetAll() ([]Project, error) {
// 	projects := []Projects{}
//
// 	dataDir, err := server.NewDataDir()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	wwwPath := dataDir.GetWWWPath()
// 	ProjectPath := dataDir.GetProjectPath()
//
// 	err = filepath.Walk(ProjectPath, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			log.Panic("Error Walking %s: %v", ProjectPath, err)
// 		}
// 		if !info.IsDir() {
// 			log.Panic("Directory not found at %s: %v", ProjectPath, err)
// 		}
// 		if path != ProjectPath {
// 			_, projectName := filepath.Split()
// 		}
// 	})
// }

// TODO: SOURCE and create projects for configs manually added to config path
// and create appropriate directories.

// func (p *Project) createConfig() error {
//
// }

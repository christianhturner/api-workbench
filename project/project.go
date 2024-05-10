package project

import (
	"log"
	"os"
	"path/filepath"

	"github.com/christianhturner/api-workbench/server"
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
	name       string
	wwwPath    string
	configPath string
}

func New(name string) (*Project, error) {
	dataDir, err := server.NewDataDir()
	if err != nil {
		log.Panic(err)
	}
	// TODO: ~Create wwwPath project subdirectory~
	// - Test this works
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
	// TODO: Create configPath project subdirectory
	// - Test this works
	projectConfigPath := filepath.Join(dataDir.GetConfigPath(), name)
	_, err = os.Stat(projectConfigPath)
	if os.IsNotExist(err) {
		log.Printf("Creating config sub directory for %s at %s", name, projectConfigPath)
		err = os.Mkdir(projectConfigPath, 0755)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Path, %s already exist for project %s.", projectConfigPath, name)
	}
	return &Project{
		name:       name,
		wwwPath:    dataDir.GetWWWPath(),
		configPath: dataDir.GetConfigPath(),
	}, nil
}

// TODO: SOURCE and create projects for configs manually added to config path
// and create appropriate directories.

// func (p *Project) createConfig() error {
//
// }

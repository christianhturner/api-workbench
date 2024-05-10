package server

import (
	"log"
	"os"
	"path/filepath"
)

type DataDir struct {
	rootPath    string
	wwwPath     string
	configpath  string
	projectPath string
}

func NewDataDir() (*DataDir, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	rootPath := filepath.Join(homeDir, ".local", "share", "api-workbench")
	wwwPath := filepath.Join(rootPath, "www")
	configPath := filepath.Join(rootPath, "config.json")
	projectPath := filepath.Join(rootPath, "projects")

	return &DataDir{
		rootPath:    rootPath,
		wwwPath:     wwwPath,
		configpath:  configPath,
		projectPath: projectPath,
	}, nil
}

func (dd *DataDir) CreateDataDirs() error {
	for _, dir := range []string{dd.rootPath, dd.wwwPath, dd.projectPath} {
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			log.Printf("Creating directory at %s", dir)
			err = os.Mkdir(dir, 0755)
			if err != nil {
				return err
			}
		}
	}

	_, err := os.Stat(dd.configpath)
	if os.IsNotExist(err) {
		log.Printf("Creating config file at %s", dd.configpath)
		_, err = os.Create(dd.configpath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dd *DataDir) GetWWWPath() string {
	return dd.wwwPath
}

func (dd *DataDir) GetRootPath() string {
	return dd.rootPath
}

func (dd *DataDir) GetProjectPath() string {
	return dd.projectPath
}

func (dd *DataDir) GetConfigPath() string {
	return dd.configpath
}

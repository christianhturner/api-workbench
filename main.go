package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/christianhturner/api-workbench/project"
	"github.com/christianhturner/api-workbench/server"
	mainMenu "github.com/christianhturner/api-workbench/tui/mainmenu"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	logFile, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("Fatal...", err)
		os.Exit(1)
	}
	defer logFile.Close()
	// createNetPath()
	dd, err := server.NewDataDir()
	if err != nil {
		log.Fatal(err)
	}
	err = dd.CreateDataDirs()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("root: %s\nwww: %s\nprojects: %s", dd.GetRootPath(), dd.GetWWWPath(), dd.GetProjectPath())
	db, err := openSqlite()
	if err != nil {
		log.Fatal(err)
	}
	pr := project.DB{DB: db}
	projects, err := pr.GetAllProjects()
	if err != nil {
		log.Fatal(err)
	}
	if len(projects) < 1 {
		name := project.NewProjectPrompt()
		_, err := pr.CreateDBEntry(name)
		if err != nil {
			log.Fatalf("error creating project: %v", err)
		}
		p := tea.NewProgram(mainMenu.InitialModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Printf("Alas, there has been an error: %v", err)
			os.Exit(1)
		}
	} else {

		p := tea.NewProgram(mainMenu.InitialModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}
}

func createNetPath() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dataDir := filepath.Join(homeDir, ".local", "share", "api-workbench", "www")

	_, err = os.Stat(dataDir)
	if os.IsNotExist(err) {
		log.Printf("Creating data directory at %s", dataDir)
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			log.Panic(err)
		}
	}
}

func openSqlite() (*gorm.DB, error) {
	dataDir, err := server.NewDataDir()
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(sqlite.Open(dataDir.GetRootPath()+"/"+"api-workbench.db"), &gorm.Config{})
	if err != nil {
		log.Printf("cannot open DB for: %v", err)
		return db, fmt.Errorf("cannot open DB for: %v", err)
	}
	err = db.AutoMigrate(&project.Project{})
	if err != nil {
		log.Printf("unable to migrate database: %v", err)
		return db, fmt.Errorf("unable to migrate database: %w", err)
	}
	return db, nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/christianhturner/api-workbench/server"
	"github.com/christianhturner/api-workbench/tui/mainMenu"
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

	p := tea.NewProgram(mainMenu.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
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

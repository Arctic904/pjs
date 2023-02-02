package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/arctic904/pjs/tui/constants"
	"github.com/arctic904/pjs/utils"
	tea "github.com/charmbracelet/bubbletea"
)

// StartTea the entry point for the UI. Initializes the model.
func StartTea(pr []utils.Project) {
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		log.Fatalln("Couldn't open a file for logging:", err)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	constants.Pr = &pr

	m := InitProject()
	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if err := constants.P.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

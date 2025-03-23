package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Println("Hello World!")

	p := tea.NewProgram(model, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error starting program:", err)
	}
}

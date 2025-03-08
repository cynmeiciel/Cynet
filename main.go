// cmd/wifi-manager/main.go
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	selected string
	items    []string
}

func initialModel() model {
	networks, err := fetchWiFis()

	if err != nil {
		networks = []string{"Error!"}
	}

	return model{
		cursor: 0,
		items:  networks,
	}
}

func (m model) View() string {
	s := "Cynet - iwctl Helper\n\n"

	for i, item := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, item)
	}

	if m.selected != "" {
		s += fmt.Sprintf("\nSelected: %s\n", m.selected)
	}

	s += "\nPress Esc or q to quit, h to show help\n"
	return s
}

func (m model) Init() tea.Cmd {
	scanWiFis()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.items) - 1
			}
		case "down":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "right":
			m.selected = m.items[m.cursor]
		}
	}

	return m, nil
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

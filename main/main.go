package main

import (
	// "github.com/ANewJanuary/Stattracker/funcs"
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
		Blink(true).
    PaddingTop(2).
    PaddingLeft(4).
    Width(22)

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"Birth", "Death", "Rebirth"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 { m.cursor-- }

		case "down", "j":
			if m.cursor > len(m.choices)-1 {m.cursor++}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok { delete(m.selected, m.cursor) 
		} else { m.selected[m.cursor] = struct{}{} }
	}
}
		return m, nil
}

func (m model) View() string {
	s := "Current Level:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i { cursor = ">" }

		checked := " "
		if _, ok := m.selected[i]; ok {checked = "x"}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

		s += "\n Press q to quit. \n"

		return s
}


func main() {
	fmt.Println(style.Render("Hello, kitty"))
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

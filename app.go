package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	doc := strings.Builder{}
	// Tabs
	{
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			activeTab.Render("Notes"),
			tab.Render("TODO"),
			tab.Render("Focus"),
		)
		gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		doc.WriteString(row + "")
	}
	fmt.Println(docStyle.Render(doc.String()))
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

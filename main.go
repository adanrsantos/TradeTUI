package main

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
	"github.com/adanrsantos/TradeTUI/app"
	"os"
)

func main() {
	model := app.New()

	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

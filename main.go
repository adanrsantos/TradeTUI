package main

import (
	"fmt"
	"os"

	"github.com/adanrsantos/TradeTUI/internals/app"
	"github.com/adanrsantos/TradeTUI/internals/providers/databento"

	tea "charm.land/bubbletea/v2"
)

func main() {
	provider := databento.New()

	model := app.New(provider)

	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

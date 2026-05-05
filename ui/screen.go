package ui

import (
	"github.com/adanrsantos/TradeTUI/types"
)

const (
	MainMenuView types.MenuScreen = iota
	TimeFrameMenuView
	SymbolMenuView
	StartDateMenuView
	EndDateMenuView
	LimitMenuView
)

const (
	MainMenuFocus types.Focus = iota
	SubmitFocus
)

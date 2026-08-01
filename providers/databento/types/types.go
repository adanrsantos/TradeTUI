package types

import (
	"github.com/adanrsantos/TradeTUI/types"
	"time"
)

type DatabentoModel struct {
	MainCursor    int
	CursorStack   []int
	SubmitCursor  int
	SettingCursor int
	Query         Query
	Screen        Screen
	Focus         Focus
	Cfg           *types.ProviderDetails
	Secret        string
}

type Query struct {
	Dataset   *Dataset
	Symbol    *Symbol
	Schema    *Schema
	StartDate time.Time
	EndDate   time.Time
	Limit     int
}

type Dataset struct {
	Display string
	Value   string
	Symbols []Symbol
}

type Symbol struct {
	Display string
	Value   string
}

type Schema struct {
	Display string
	Value   string
}

type MenuItem struct {
	Label  string
	Target Screen
	Action Action
}

type Action int

const (
	SubmitAction Action = iota
	ResetAction
)

type Focus int

const (
	MainFocus Focus = iota
	QuerySubmitFocus
	SubmitFocus
	SettingFocus
)

type Screen int

const (
	MainMenuScreen Screen = iota
	HistMenuScreen
	StatScreen
	DownloadScreen
	SettingScreen

	HistDataset
	HistSymbol
	HistSchema
	HistStart
	HistEnd
	HistLimit
	HistRequest
)

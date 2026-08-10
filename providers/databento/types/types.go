package types

import (
	"charm.land/bubbles/v2/textinput"
	"github.com/adanrsantos/TradeTUI/types"
	"time"
)

type DatabentoModel struct {
	MainCursor    int
	CursorStack   []int
	SubmitCursor  int
	SettingCursor int
	Input         textinput.Model
	ErrInput      error
	Mode          Mode
	Query         Query
	ErrQuery      error
	Screen        Screen
	Focus         Focus
	Cfg           *types.ProviderDetails
	Secret        string
}

type Query struct {
	Dataset   *Dataset
	Symbol    *Symbol
	Schema    *Schema
	StartDate *time.Time
	EndDate   *time.Time
	Limit     *int
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
	Field  QueryField
	Action Action
}

type TimeFormat struct {
	Layout      string
	Placeholder string
	CharLimit   int
}

type Mode int

const (
	NormalMode Mode = iota
	EditMode
)

type QueryField int

const (
	DatasetField QueryField = iota
	SymbolField
	SchemaField
	StartField
	EndField
	LimitField
)

type Action int

const (
	SubmitAction Action = iota
	ResetAction
	ContinueAction
	CancelAction
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

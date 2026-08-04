package types

import (
	"charm.land/bubbles/v2/textinput"
	"github.com/adanrsantos/TradeTUI/types"
)

type DatabentoModel struct {
	MainCursor    int
	CursorStack   []int
	SubmitCursor  int
	SettingCursor int
	TextInput     textinput.Model
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
	StartDate *TimePreset
	EndDate   *TimePreset
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

type TimeValue string

const (
	MarketOpen  TimeValue = "09:30"
	MarketClose TimeValue = "16:00"
	AsiaOpen    TimeValue = "20:00"
	LondonOpen  TimeValue = "03:00"
	Midnight    TimeValue = "00:00"
	Noon        TimeValue = "12:00"
)

type TimePreset struct {
	Display string
	Value   TimeValue
}

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

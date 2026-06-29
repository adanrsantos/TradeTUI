package databento

import "github.com/adanrsantos/TradeTUI/types"

type databentoModel struct {
	settingCursor int
	mainCursor    int
	screen        Screen
	focus         Focus
	cfg           *types.ProviderDetails
}

type Focus int

const (
	MainFocus Focus = iota
	SettingFocus
)

type Screen int

const (
	MainMenuScreen Screen = iota
	HistoricalScreen
	LiveScreen
	TestDataScreen
	SymbolScreen
	SchemaScreen
	StartScreen
	EndScreen
	LimitScreen
)

type MenuItem struct {
	Label  string
	Target Screen
}

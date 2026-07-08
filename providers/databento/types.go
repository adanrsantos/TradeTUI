package databento

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

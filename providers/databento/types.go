package databento

type Focus int

const (
	MainFocus Focus = iota
	SettingFocus
)

type Screen int

const (
	MainMenuScreen Screen = iota

	HistMenuScreen
	HistSymbol
	HistSchema
	HistStart
	HistEnd
	HistLimit

	LiveMenuScreen
	LiveSymbol
	LiveSchema
)

type MenuItem struct {
	Label  string
	Target Screen
}

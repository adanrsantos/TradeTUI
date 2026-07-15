package databento

import "time"

type Query struct {
	TimeFrame TimeFrame
	Symbol    Symbol
	StartDate time.Time
	EndDate   time.Time
	Limit     int
}

type TimeFrame string

const (
	OneSecond     TimeFrame = "1sec"
	OneMinute     TimeFrame = "1min"
	FifteenMinute TimeFrame = "15min"
	OneHour       TimeFrame = "1hour"
	FourHour      TimeFrame = "4hour"
	Daily         TimeFrame = "daily"
)

type Symbol string

const (
	NQ Symbol = "NQ"
	ES Symbol = "ES"
	YM Symbol = "YM"
)

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

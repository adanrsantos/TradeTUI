package databento

import "time"

type Query struct {
	Symbol    Symbol
	Schema    Schema
	StartDate time.Time
	EndDate   time.Time
	Limit     int
}

type Schema string

const (
	OneSecond Schema = "ohlcv-1s"
	OneMinute Schema = "ohlcv-1m"
	OneHour   Schema = "ohlcv-1h"
	Daily     Schema = "ohlcv-1d"
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

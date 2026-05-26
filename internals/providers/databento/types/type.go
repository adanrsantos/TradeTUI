package types

import (
	"time"
)

type Model struct {
	Screen         MenuScreen
	Config         QueryConfig
	Focus          Focus
	History        []HistoryItem
	MainMenuCursor int
	SubmitCursor   int
	Err            string
}

type QueryConfig struct {
	TimeFrame TimeFrame
	Symbol    Symbol
	StartDate time.Time
	EndDate   time.Time
	Limit     int
}

type MenuItem struct {
	Label  string
	Target MenuScreen
}

type HistoryItem struct {
	Config    QueryConfig
	Timestamp time.Time
}

type DatePreset struct {
	Label string
	Value func() time.Time
}

type MenuScreen int
type Focus int

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
)

type Limit int

const (
	Zero Limit = 0
	Ten Limit = 10
	OneHundred Limit = 100
	OneThousand Limit = 1000
)

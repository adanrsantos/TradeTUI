package types

import "time"

type QueryConfig struct {
	TimeFrame TimeFrame
	Symbol Symbol
	StartDate time.Time
	EndDate time.Time
	Limit int
}

type TimeFrame string

const (
	OneSecond TimeFrame = "1sec"
	OneMinute TimeFrame = "1min"
	FifteenMinute TimeFrame = "15min"
	OneHour TimeFrame = "1hour"
	FourHour TimeFrame = "4hour"
	Daily TimeFrame = "daily"
)

type Symbol string

const (
	NQ Symbol = "NQ"
	ES Symbol = "ES"
)
package databento

import "github.com/adanrsantos/TradeTUI/types"

type databentoModel struct {
	settingCursor int
	mainCursor    int
	focus         focus
	cfg           *types.ProviderDetails
}

type focus int

const (
	MainFocus focus = iota
	SettingFocus
)

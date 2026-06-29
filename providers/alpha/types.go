package alpha

import "github.com/adanrsantos/TradeTUI/types"

type alphaModel struct {
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

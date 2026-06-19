package databento

type databentoModel struct {
	settingCursor int
	mainCursor    int
	focus         focus
}

type focus int

const (
	MainFocus focus = iota
	SettingFocus
)

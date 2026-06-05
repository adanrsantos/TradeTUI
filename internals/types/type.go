package types

type Model struct {
	Focus          FocusView
	MainMenuCursor int
	SubmitCursor   int
	SettingCursor  int
	Err            string
}

type FocusView int

type MenuItem struct {
	Label  string
	Target string
}

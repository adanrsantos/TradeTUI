package types

type Model struct {
	Screen        Screen
	SettingCursor int
	Cfg           *Config
	StateFilePath string
}

type Screen int

const (
	SettingScreen Screen = iota
	ProviderScreen
)

type ProviderDetails struct {
	APIKey   string `json:"api_key,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
}

type Config struct {
	User struct {
		Theme           string `json:"theme"`
		Language        string `json:"language"`
		CurrentProvider string `json:"current_provider"`
	}
	Provider struct {
		Databento *ProviderDetails `json:"databento,omitempty"`
		Alpha     *ProviderDetails `json:"alpha,omitempty"`
	}
}

package app

import "github.com/adanrsantos/TradeTUI/internals/providers"

func New(p providers.Provider) Model {
	return Model{
		width: 30,
		height: 30,
		Provider: p,
	}
}
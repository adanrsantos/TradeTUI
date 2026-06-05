package app

import "github.com/adanrsantos/TradeTUI/internals/providers"

func New(p providers.Provider) Model {
	return Model{
		Provider: p,
	}
}


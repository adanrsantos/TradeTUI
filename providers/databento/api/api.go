package api

import (
	"fmt"
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
)

func FetchHistory(query types.Query) error {
	fmt.Println(query)
	return nil
}

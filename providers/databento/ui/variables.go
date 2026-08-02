package ui

import (
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
)

var Parent = map[types.Screen]types.Screen{
	types.HistMenuScreen: types.MainMenuScreen,
	types.StatScreen:     types.MainMenuScreen,
	types.DownloadScreen: types.MainMenuScreen,
	types.SettingScreen:  types.MainMenuScreen,

	types.HistDataset: types.HistMenuScreen,
	types.HistSymbol:  types.HistMenuScreen,
	types.HistSchema:  types.HistMenuScreen,
	types.HistStart:   types.HistMenuScreen,
	types.HistEnd:     types.HistMenuScreen,
	types.HistLimit:   types.HistMenuScreen,
}

var Menu = []types.MenuItem{
	{Label: "Historical", Target: types.HistMenuScreen},
	{Label: "Statistics", Target: types.StatScreen},
	{Label: "Downloads", Target: types.DownloadScreen},
	{Label: "Settings", Target: types.SettingScreen},
}

var HistoricalMenu = []types.MenuItem{
	{Label: "Dataset", Target: types.HistDataset, Field: types.DatasetField},
	{Label: "Symbol", Target: types.HistSymbol, Field: types.SymbolField},
	{Label: "Schema", Target: types.HistSchema, Field: types.SchemaField},
	{Label: "Start", Target: types.HistStart, Field: types.StartField},
	{Label: "End", Target: types.HistEnd, Field: types.EndField},
	{Label: "Limit", Target: types.HistLimit, Field: types.LimitField},
}

var SubmitChoices = []types.MenuItem{
	{Label: "Submit", Action: types.SubmitAction},
	{Label: "Reset", Action: types.ResetAction},
}

var Datasets = []types.Dataset{
	{
		Display: "CME Globex MDP 3.0",
		Value:   "GLBX.MDP3",
		Symbols: FutureSymbols,
	},
}

var FutureSymbols = []types.Symbol{
	{
		Display: "E-mini Nasdaq-100 (NQ)",
		Value:   "NQ",
	},
	{
		Display: "E-mini S&P 500 (ES)",
		Value:   "ES",
	},
	{
		Display: "E-mini Dow Jones (YM)",
		Value:   "YM",
	},
}

var Schemas = []types.Schema{
	{
		Display: "1 second (ohlcv-1s)",
		Value:   "ohlcv-1s",
	},
	{
		Display: "1 minute (ohlcv-1m)",
		Value:   "ohlcv-1s",
	},
	{
		Display: "1 hour (ohlcv-1h)",
		Value:   "ohlcv-1s",
	},
	{
		Display: "daily (ohlcv-1d)",
		Value:   "ohlcv-1s",
	},
}

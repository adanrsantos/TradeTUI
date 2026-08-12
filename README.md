# TradeTUI

TradeTUI is a terminal-based trading application written in Go. It provides a TUI for querying and working with historical market data, with a provider-based architecture designed to support multiple market-data providers.

## Features

- Terminal-based trading interface
- Provider-based architecture for market-data integrations
- Historical market-data requests
- Databento historical data integration
- Historical request cost estimation through Databento
- Configurable dataset, symbol, schema, start time, end time, and record limit
- New York timezone-aware date/time input
- Query validation before making requests
- JSON/JSON Lines-friendly candle storage
- Environment-based API key configuration
- Keyboard-driven navigation

## Tech Stack
- Go
- Bubble Tea — TUI framework
- Lip Gloss — terminal styling
- Databento — historical market data
- JSON / JSON Lines — market-data serialization
- godotenv — environment configuration

## Getting Started
### Requirements
- Go 1.24+ recommended
- A Databento API key for Databento functionality
### Clone the repository
git clone https://github.com/adanrsantos/TradeTUI.git
cd TradeTUI
### Configure environment variables
Create a .env file in the project root:
DATABENTO_API_KEY=your_api_key_here
API keys are intentionally kept out of application state/configuration files.
### Run TradeTUI
go run main.go

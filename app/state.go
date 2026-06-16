package app

import (
	"encoding/json"
	"fmt"
	"github.com/adanrsantos/TradeTUI/types"
	"os"
	"path/filepath"
)

func findStateFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := cwd
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			path := filepath.Join(dir, "state.json")

			return path, nil
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("could not find project root (missing go.mod)")
}

func loadConfig(path string, config *types.Config) error {
	fileData, err := os.ReadFile(path)
	if err == nil {
		return json.Unmarshal(fileData, config)
	}
	if !os.IsNotExist(err) {
		return err
	}

	config.User.Theme = "dark"
	config.User.Language = "en-US"

	return saveConfig(path, config)
}

func saveConfig(path string, config *types.Config) error {
	jsonData, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}

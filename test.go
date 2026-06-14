package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func findProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := cwd
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("could not find project root (missing go.mod)")
}

func openStateFile(path string) (*os.File, error) {
	stateFilePath := filepath.Join(path, "state.json")

	f, err := os.OpenFile(stateFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

type ProviderDetails struct {
	APIKey   string `json:"api_key"`
	Endpoint string `json:"endpoint"`
}

type Config struct {
	User struct {
		Theme    string `json:"theme"`
		Language string `json:"language"`
	} `json:"user"`
	Providers struct {
		Databento ProviderDetails `json:"databento"`
		Alpha     ProviderDetails `json:"alpha"`
	} `json:"providers"`
}

func saveDummyConfig(root string) error {
	cfg := Config{}

	cfg.User.Theme = "dark"
	cfg.User.Language = "en-US"

	cfg.Providers.Databento = ProviderDetails{
		APIKey:   "db_123456",
		Endpoint: "https://hist.databento.com",
	}
	cfg.Providers.Alpha = ProviderDetails{
		APIKey:   "alpha_vantage_789",
		Endpoint: "https://www.alphavantage.co",
	}

	jsonData, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	path := filepath.Join(root, "state.json")

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func test() {
	projectRoot, err := findProjectRoot()
	if err != nil {
		log.Fatalf("Error locating project: %v", err)
	}

	err = saveDummyConfig(projectRoot)
	if err != nil {
		log.Fatalf("Error saving dummy data: %v", err)
	}

	file, err := openStateFile(projectRoot)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
}

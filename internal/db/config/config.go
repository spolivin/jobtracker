package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/term"
)

// Database connection config.
type ConnectionConfig struct {
	DBHost string `json:"db_host"`
	DBPort int    `json:"db_port"`
	DBUser string `json:"db_user"`
	DBName string `json:"db_name"`
}

// get_config_path retrieves a path to database connection config.
func get_config_path() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "jobtracker", "config.json"), nil
}

// LoadConfig loads config for connection to the database.
func LoadConfig() (*ConnectionConfig, error) {
	p, err := get_config_path()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	var c ConnectionConfig
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

// SaveConfig saves a new config to a default path.
func SaveConfig(c *ConnectionConfig) (string, error) {
	p, err := get_config_path()
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Dir(p), 0700); err != nil {
		return "", err
	}

	data, _ := json.MarshalIndent(c, "", "  ")
	return p, os.WriteFile(p, data, 0600)
}

// PromptPassword prompts for a user to enter password to Postgres.
func PromptPassword() (string, error) {
	fmt.Print("Postgres password: ")
	b, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return string(b), err
}

package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)


const confName = ".gatorconfig.json"

type Config struct {
	DbURL		string	`json:"db_url"`
	CurrentUser	string	`json:"current_user_name"`
}

func Read() (Config, error) {
	var r Config

	confPath, err := getConfigFilePath()

	if err != nil {
		return r, err
	}

	dat, err := os.ReadFile(confPath)

	if err != nil {
		return r, err 
	}

	err = json.Unmarshal(dat, &r)

	if err != nil {
		return r, err
	}

	return r, nil
}

func write(cfg Config) error {
	dat, err := json.MarshalIndent(cfg, "", "  ")

	if err != nil {
		return err
	}

	confPath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	err = os.WriteFile(confPath, dat, 0644)

	return err
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	confPath := filepath.Join(home, confName)
	return confPath, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUser = name
	err := write(*cfg)

	return err
}

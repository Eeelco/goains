package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func initializeConfigVariables() {
	HOME_DIR, _ := os.UserHomeDir()
	CONFIG_DIR = filepath.Join(HOME_DIR, ".config/"+APP_NAME)
	CONFIG_FILE = filepath.Join(CONFIG_DIR, "config.json")
	DB_FILE = filepath.Join("exercises.json")
}

func configFolderExists() bool {
	_, err := os.Stat(CONFIG_DIR)
	return !os.IsNotExist(err)
}

func createConfigFolder() {
	os.Mkdir(CONFIG_DIR, 0755)
	os.Mkdir(CONFIG_DIR+"/plans", 0755)
	var default_cfg = DefaultCfg()

	config = default_cfg

	f, _ := os.Create(CONFIG_FILE)
	defer f.Close()
	cfg_json, _ := json.MarshalIndent(default_cfg, "", "  ")
	f.Write(cfg_json)
}

func LoadConfigAndDB() {
	f, _ := os.Open(CONFIG_FILE)
	defer f.Close()
	dec := json.NewDecoder(f)
	_ = dec.Decode(&config)

	ex, _ := os.Open(DB_FILE)
	defer ex.Close()
	dec = json.NewDecoder(ex)
	_ = dec.Decode(&exerciseDatabase)
}

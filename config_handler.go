package main

import (
	"encoding/json"
	"log"
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

func saveDefaultPlans() {
	plan_dir := CONFIG_DIR + "/plans"
	plan_files, _ := os.ReadDir("default_plans")
	for _, plan := range plan_files {
		filebytes, err := os.ReadFile("default_plans/" + plan.Name())
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(plan_dir+"/"+plan.Name(), filebytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
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

	saveDefaultPlans()
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

package main

import (
	"embed"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
)

// This file contains functions related to handling configuration and data files.

// initializeConfigVariables initializes the configuration variables such as the home directory and the configuration directory.
func initializeConfigVariables() {
	HOME_DIR, _ := os.UserHomeDir()
	CONFIG_DIR = filepath.Join(HOME_DIR, ".config", APP_NAME)
	CONFIG_FILE = filepath.Join(CONFIG_DIR, "config.json")
}

// FolderExists checks if the given folder exists.
func FolderExists(folder string) bool {
	_, err := os.Stat(folder)
	return !os.IsNotExist(err)
}

//go:embed default_plans
var default_plan_folder embed.FS

func saveDefaultPlans() {
	plan_dir := CONFIG_DIR + "/plans"

	plan_files, _ := default_plan_folder.ReadDir("default_plans")
	for _, plan := range plan_files {
		filebytes, err := default_plan_folder.ReadFile("default_plans/" + plan.Name())
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(plan_dir+"/"+plan.Name(), filebytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// createConfigFolder creates the configuration folder and saves the default configuration.
func createConfigFolder() {
	os.Mkdir(CONFIG_DIR, 0755)
	os.Mkdir(CONFIG_DIR+"/plans", 0755)
	os.Mkdir(CONFIG_DIR+"/workouts", 0755)
	var default_cfg = DefaultCfg()

	config = default_cfg

	f, _ := os.Create(CONFIG_FILE)
	defer f.Close()
	cfg_json, _ := json.MarshalIndent(default_cfg, "", "  ")
	f.Write(cfg_json)

	saveDefaultPlans()
}

//go:embed exercises.json
var db_file embed.FS

func LoadConfigAndDB() {
	f, _ := os.Open(CONFIG_FILE)
	defer f.Close()
	dec := json.NewDecoder(f)
	_ = dec.Decode(&config)

	data, _ := db_file.ReadFile("exercises.json")

	json.Unmarshal(data, &exerciseDatabase)
}

func saveWorkout(data PlanDay) {
	end := time.Now()
	duration := end.Sub(start_time)
	filename := CONFIG_DIR + "/workouts/" + config.CurrentPlan + ".json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(CONFIG_DIR + "/workouts/" + data.Name + ".json")
	}
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	with_metadata := Progress{
		Start:    start_time.String(),
		Duration: duration.Seconds(),
		Data:     data,
	}

	json_data, err := json.MarshalIndent(with_metadata, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	f.Write(json_data)
	f.Write(([]byte)(",\n"))
}

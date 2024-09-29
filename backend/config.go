package backend

import (
	"embed"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
)

// This file contains functions related to handling app configuration.
// This includes loading and saving configuration settings, plans, and last workouts.

// Config represents the configuration settings for the application.
type Config struct {
	CurrentPlan   string // Identifier of the current plan
	NextDayIdx    int    // Index of the next day in the plan
	DefaultNrSets int    // Default number of sets for an exercise
	DefaultNrReps int    // Default number of repetitions for an exercise
	DefaultRest   int    // Default rest time in seconds between exercises
	UserHomeDir   string // User's home directory
	ConfigDir     string // Configuration directory
	ConfigFile    string // Configuration file
}

// DefaultCfg returns a Config instance with default values.
func DefaultCfg() Config {
	return Config{
		CurrentPlan:   "",
		NextDayIdx:    0,
		DefaultNrSets: 3,
		DefaultNrReps: 10,
		DefaultRest:   60,
		UserHomeDir:   "",
		ConfigDir:     "",
		ConfigFile:    "",
	}
}

// Initialize sets up the configuration directories and loads previously saved configurations
func (c *Config) Initialize() {
	c.UserHomeDir, _ = os.UserHomeDir()
	c.ConfigDir = filepath.Join(c.UserHomeDir, ".config", APP_NAME)
	c.ConfigFile = filepath.Join(c.ConfigDir, "config.json")

	if !FolderExists(c.ConfigDir) {
		c.createConfigFolder()
	}
	c.loadFile()
}

// createConfigFolder creates the configuration folder and saves the default configuration.
func (c *Config) createConfigFolder() {
	os.Mkdir(c.ConfigDir, 0755)
	os.Mkdir(c.ConfigDir+"/plans", 0755)
	os.Mkdir(c.ConfigDir+"/last_workouts", 0755)

	*c = DefaultCfg()

	f, _ := os.Create(c.ConfigFile)
	defer f.Close()
	cfg_json, _ := json.MarshalIndent(c, "", "  ")
	f.Write(cfg_json)

	c.saveDefaultPlans()
}

// loadFile loads previously set config values from a file
func (c *Config) loadFile() {
	f, _ := os.Open(c.ConfigFile)
	defer f.Close()
	dec := json.NewDecoder(f)
	_ = dec.Decode(&c)

}

// FolderExists checks if the given folder exists.
func FolderExists(folder string) bool {
	_, err := os.Stat(folder)
	return !os.IsNotExist(err)
}

// Embedded directory containing some pre-defined workout plans
//
//go:embed default_plans
var default_plan_folder embed.FS

// saveDefaultPlans saves the embedded workout plans to the config directory
func (c Config) saveDefaultPlans() {
	plan_dir := c.ConfigDir + "/plans"

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

// saveWorkout saves the repetitions and weights from the given workout to a JSON file
func (c Config) saveWorkout(data PlanDay, start_time time.Time) {
	end := time.Now()
	duration := end.Sub(start_time)
	filename := c.ConfigDir + "/last_workouts/" + c.CurrentPlan + "_" + data.Name + ".json"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
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
		return
	}
	f.Write(json_data)
}

// GetAllPlans returns all workout plans stored in the config folder
func (c Config) GetAllPlans() []Plan {
	files, _ := os.ReadDir(c.ConfigDir + "/plans")
	plans := []Plan{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, _ := os.Open(c.ConfigDir + "/plans/" + file.Name())
		plan := Plan{}
		json.NewDecoder(f).Decode(&plan)
		plans = append(plans, plan)
	}
	return plans
}

// SaveConfig overwrites the value of c and saves the new config to disk
func (c *Config) SaveConfig(new_c Config) {
	*c = new_c
	f, _ := os.Create(c.ConfigDir)
	defer f.Close()
	cfg_json, _ := json.MarshalIndent(c, "", "  ")
	f.Write(cfg_json)
}

// GetPlan returns the workout plan with the given name
func (c Config) GetPlan(name string) Plan {
	plan := Plan{}
	f, err := os.Open(c.ConfigDir + "/plans/" + name + ".json")
	if err != nil {
		return plan
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&plan)
	return plan
}

// GetLastWorkout returns the last workout progress for the given day
func (c Config) GetLastWorkout(day string) Progress {
	filename := c.ConfigDir + "/last_workouts/" + c.CurrentPlan + "_" + day + ".json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return Progress{}
	}

	f, err := os.Open(c.ConfigDir + "/last_workouts/" + c.CurrentPlan + "_" + day + ".json")
	if err != nil {
		log.Fatal(err)
		return Progress{}
	}
	defer f.Close()
	progress := Progress{}
	json.NewDecoder(f).Decode(&progress)
	return progress
}

// SavePlan saves the given plan to a JSON file.
// It returns true if the plan is successfully saved, false otherwise.
func (c Config) SavePlan(plan Plan) bool {
	_, err := os.Stat(plan.Name)
	if !os.IsNotExist(err) {
		return false
	}

	f, err := os.Create(c.ConfigDir + "/plans/" + plan.Name + ".json")
	if err != nil {
		return false
	}

	defer f.Close()
	planJSON, _ := json.MarshalIndent(plan, "", "  ")
	f.Write(planJSON)
	return true
}

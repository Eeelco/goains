package main

import (
	"context"
	"encoding/json"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	initializeConfigVariables()
	if !configFolderExists() {
		createConfigFolder()
	}
	LoadConfigAndDB()

}

func (a *App) GetPlan() string {
	return config.CurrentPlan
}

func (a *App) GetAllPlans() []Plan {
	files, _ := os.ReadDir(CONFIG_DIR + "/plans")
	plans := []Plan{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, _ := os.Open(CONFIG_DIR + "/plans/" + file.Name())
		plan := Plan{}
		json.NewDecoder(f).Decode(&plan)
		plans = append(plans, plan)
	}
	return plans
}

func (a *App) SetCurrentPlan(plan_name string) {
	config.CurrentPlan = plan_name

	f, _ := os.Create(CONFIG_FILE)
	defer f.Close()
	cfg_json, _ := json.MarshalIndent(config, "", "  ")
	f.Write(cfg_json)
}

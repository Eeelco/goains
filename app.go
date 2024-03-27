package main

import (
	"context"
	"encoding/json"
	"os"
	"strings"
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

func (a *App) HasPlan() bool {
	return config.CurrentPlan != ""
}

func (a *App) GetImgPrefix() string {
	return CONFIG_DIR + "/img/"
}

func (a *App) GetExercises(filter string) []Exercise {
	if filter == "" {
		return exerciseDatabase
	}
	out := []Exercise{}
	for _, ex := range exerciseDatabase {
		if strings.Contains(strings.ToLower(ex.Name), strings.ToLower(filter)) {
			out = append(out, ex)
		}
	}

	if MAX_NR_EXERCISES < len(out) {
		return out[:MAX_NR_EXERCISES]
	} else {
		return out
	}
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

func (a *App) GetConfig() Config {
	return config
}

func (a *App) SavePlan(plan Plan) {
	f, err := os.Create(CONFIG_DIR + "/plans/" + plan.Name + ".json")
	if err != nil {
		i := 0
		new_file := CONFIG_DIR + "/plans/" + plan.Name + "_" + string(rune(i)) + ".json"
		for {
			if _, err := os.Stat(new_file); os.IsNotExist(err) {
				f, _ = os.Create(new_file)
				break
			}
			i++
			new_file = CONFIG_DIR + "/plans/" + plan.Name + "_" + string(rune(i)) + ".json"
		}
	}
	defer f.Close()
	plan_json, _ := json.MarshalIndent(plan, "", "  ")
	f.Write(plan_json)
}

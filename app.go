package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	if !FolderExists(CONFIG_DIR) {
		createConfigFolder()
	}
	LoadConfigAndDB()

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

func (a *App) SaveConfig(c Config) {
	config = c
	f, _ := os.Create(CONFIG_FILE)
	defer f.Close()
	cfg_json, _ := json.MarshalIndent(c, "", "  ")
	f.Write(cfg_json)
}

func (a *App) GetPlan(name string) Plan {
	f, _ := os.Open(CONFIG_DIR + "/plans/" + name + ".json")
	defer f.Close()
	plan := Plan{}
	json.NewDecoder(f).Decode(&plan)
	return plan
}

func (a *App) GetExerciseById(id string) Exercise {
	for _, ex := range exerciseDatabase {
		if ex.Id == id {
			return ex
		}
	}
	return Exercise{}
}

func (a *App) GetExercisesByIDs(ids []string) []Exercise {
	exercises := []Exercise{}
	for _, id := range ids {
		exercises = append(exercises, a.GetExerciseById(id))
	}
	return exercises
}

func (a *App) StartWorkout() {
	start_time = time.Now()
}

func (a *App) WorkoutExitDialog() bool {
	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Exit Workout",
		Message: "Are you sure you want to exit the workout?",
	})
	return result == "Yes" || result == "Ok"
}

func (a *App) WorkoutSaveDialog(data PlanDay) bool {
	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Save Workout",
		Message: "Do you want to exit and save the workout?",
	})
	if result == "Yes" || result == "Ok" {
		saveWorkout(data)
	}
	return result == "Yes" || result == "Ok"
}

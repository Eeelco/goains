package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/lithammer/fuzzysearch/fuzzy"

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
	// Set window to fullscreen on mobile
	screens, _ := runtime.ScreenGetAll(ctx)
	for _, s := range screens {
		if s.IsCurrent {
			if s.Size.Width < 1024 {
				runtime.WindowFullscreen(ctx)
			}
		}
	}
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
		if fuzzy.MatchFold(filter, ex.Name) {
			out = append(out, ex)
		}
	}
	return out
}

func (a *App) GetConfig() Config {
	return config
}

func (a *App) SavePlan(plan Plan) {
	f, err := os.Create(CONFIG_DIR + "/plans/" + plan.Name + ".json")
	if err != nil {
		i := 0
		new_file := CONFIG_DIR + "/plans/" + plan.Name + "_" + string(i) + ".json"
		for {
			if _, err := os.Stat(new_file); os.IsNotExist(err) {
				f, _ = os.Create(new_file)
				break
			}
			i++
			new_file = CONFIG_DIR + "/plans/" + plan.Name + "_" + string(i) + ".json"
		}
	}
	defer f.Close()
	plan_json, _ := json.MarshalIndent(plan, "", "  ")
	f.Write(plan_json)
}

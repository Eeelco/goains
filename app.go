package main

import (
	"context"
	"encoding/json"
	"os"

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
	if !CheckConfigFolder() {
		a.DownloadDatabase()
	} else {
		LoadConfigAndDB()
	}
}

// Check if the config folder exists, if not, create it
func CheckConfigFolder() bool {
	if !configFolderExists() {
		createConfigFolder()
		return false
	}
	return true
}

// Conditionally download exercise database
func (a *App) DownloadDatabase() {
	askDownloadOptions := runtime.MessageDialogOptions{Type: runtime.QuestionDialog, Title: "Download Database?", Message: "No exercise database found. Would you like to download it now (approx. 100 MB)?"}
	response, err := runtime.MessageDialog(a.ctx, askDownloadOptions)

	if err != nil || response != "Yes" {
		return
	}
	runtime.EventsEmit(a.ctx, "download_db")
	DownloadFile(exercises_json, DB_FILE)

	exercises, _ := os.Open(DB_FILE)
	defer exercises.Close()

	dec := json.NewDecoder(exercises)
	_ = dec.Decode(&exerciseDatabase)

	os.Mkdir(CONFIG_DIR+"/img", 0755)

	for i, ex := range exerciseDatabase {
		runtime.EventsEmit(a.ctx, "download_db_progress", [2]int{i, len(exerciseDatabase)})
		os.Mkdir(CONFIG_DIR+"/img/"+ex.Id, 0755)
		for _, img := range ex.Images {
			DownloadFile(image_url_prefix+img, CONFIG_DIR+"/img/"+img)
		}

	}
	runtime.EventsEmit(a.ctx, "download_db_done")

}

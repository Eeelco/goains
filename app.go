package main

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App is the main application struct that holds the configuration and database.
// This file contains the definition of the App struct and its methods.
// These methods can be called from the frontend using the Wails runtime.

// App struct
type App struct {
	ctx        context.Context  // The context is used to call the runtime methods
	config     Config           // Config
	db         ExerciseDatabase // Database
	start_time time.Time        // Start time of the workout
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	a.config.Initialize()
	a.db.Initialize()

}

// GetAllPlans returns a list of all plans stored in the corresponding directory.
func (a *App) GetAllPlans() []Plan {
	return a.config.GetAllPlans()
}

// SaveConfig updates the configuration settings with the given values and saves them to a JSON file.
func (a *App) SaveConfig(c Config) {
	a.config.SaveConfig(c)
}

// GetPlan returns the plan with the given name.
func (a *App) GetPlan(name string) Plan {
	return a.config.GetPlan(name)
}

// GetExercisesByIDs returns a list of exercises based on the given IDs.
// The returned list is limited to MAX_NR_EXERCISES.
func (a *App) GetExercisesByIDs(ids []string) []Exercise {
	return a.db.GetExercisesByIDs(ids)
}

// StartWorkout sets the start time of the workout.
func (a *App) StartWorkout() {
	a.start_time = time.Now()
}

// WorkoutExitDialog shows a dialog to confirm exiting the workout.
func (a *App) WorkoutExitDialog() bool {
	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Exit Workout",
		Message: "Are you sure you want to exit the workout?",
	})
	return result == "Yes" || result == "Ok"
}

// WorkoutSaveDialog shows a dialog to confirm saving the workout.
// It saves the workout progress to a JSON file if the user chooses to do so.
func (a *App) WorkoutSaveDialog(data PlanDay) bool {
	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Save Workout",
		Message: "Do you want to exit and save the workout?",
	})
	if result == "Yes" || result == "Ok" {
		a.config.saveWorkout(data, a.start_time)
	}
	return result == "Yes" || result == "Ok"
}

// GetLastWorkout returns the last workout progress for the given day.
func (a *App) GetLastWorkout(day string) Progress {
	return a.config.GetLastWorkout(day)
}

// GetExercises returns a list of exercises based on the given filter.
// If the filter is empty, it returns a subset of exerciseDatabase.
// The returned list is limited to MAX_NR_EXERCISES.
func (a *App) GetExercises(filter string) []Exercise {
	return a.db.GetExercises(filter)
}

// SavePlan saves the given plan to a JSON file.
// It returns true if the plan is successfully saved, false otherwise.
func (a *App) SavePlan(plan Plan) bool {
	return a.config.SavePlan(plan)
}

// GetConfig returns the current configuration.
func (a *App) GetConfig() Config {
	return a.config
}

// GetImgPrefix returns the image prefix used in the application.
func (a *App) GetImgPrefix() string {
	return a.config.ConfigDir + "/img/"
}

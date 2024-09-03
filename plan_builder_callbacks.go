// Package main provides a collection of functions for building plans and managing exercises.
package main

import (
	"encoding/json"
	"os"
	"strings"
)

// GetExercises returns a list of exercises based on the given filter.
// If the filter is empty, it returns a subset of exerciseDatabase.
// The returned list is limited to MAX_NR_EXERCISES.
func (a *App) GetExercises(filter string) []Exercise {
	if filter == "" {
		return exerciseDatabase[:MAX_NR_EXERCISES]
	}
	out := []Exercise{}
	for _, ex := range exerciseDatabase {
		if strings.Contains(strings.ToLower(ex.Name), strings.ToLower(filter)) {
			out = append(out, ex)
			if len(out) >= MAX_NR_EXERCISES {
				break
			}
		}
	}

	return out
}

// SavePlan saves the given plan to a JSON file.
// It returns true if the plan is successfully saved, false otherwise.
func (a *App) SavePlan(plan Plan) bool {
	_, err := os.Stat(plan.Name)
	if os.IsNotExist(err) {
		return false
	}

	f, err := os.Create(CONFIG_DIR + "/plans/" + plan.Name + ".json")
	if err != nil {
		return false
	}

	defer f.Close()
	planJSON, _ := json.MarshalIndent(plan, "", "  ")
	f.Write(planJSON)
	return true
}

// GetConfig returns the current configuration.
func (a *App) GetConfig() Config {
	return config
}

// GetImgPrefix returns the image prefix used in the application.
func (a *App) GetImgPrefix() string {
	return CONFIG_DIR + "/img/"
}

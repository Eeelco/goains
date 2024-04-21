package main

import (
	"encoding/json"
	"os"
	"strings"
)

func (a *App) GetExercises(filter string) []Exercise {
	if filter == "" {
		return exerciseDatabase[:MAX_NR_EXERCISES]
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
	plan_json, _ := json.MarshalIndent(plan, "", "  ")
	f.Write(plan_json)
	return true
}

func (a *App) GetConfig() Config {
	return config
}

func (a *App) GetImgPrefix() string {
	return CONFIG_DIR + "/img/"
}

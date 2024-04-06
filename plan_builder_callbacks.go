package main

import (
	"encoding/json"
	"os"
	"strings"
)

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

func (a *App) GetConfig() Config {
	return config
}

func (a *App) GetImgPrefix() string {
	return CONFIG_DIR + "/img/"
}

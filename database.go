package main

import (
	"embed"
	"encoding/json"
	"strings"
)

type ExerciseDatabase struct {
	Exercises []Exercise
}

//go:embed exercises.json
var db_file embed.FS

func (e *ExerciseDatabase) Initialize() {
	data, _ := db_file.ReadFile("exercises.json")

	json.Unmarshal(data, &e.Exercises)
}

func (e ExerciseDatabase) getExerciseById(id string) Exercise {
	for _, ex := range e.Exercises {
		if ex.Id == id {
			return ex
		}
	}
	return Exercise{}
}

func (e ExerciseDatabase) GetExercisesByIDs(ids []string) []Exercise {
	exercises := []Exercise{}
	for _, id := range ids {
		exercises = append(exercises, e.getExerciseById(id))
	}
	return exercises
}

// GetExercises returns a list of exercises based on the given filter.
// If the filter is empty, it returns a subset of exerciseDatabase.
// The returned list is limited to MAX_NR_EXERCISES.
func (e ExerciseDatabase) GetExercises(filter string) []Exercise {
	if filter == "" {
		return e.Exercises[:MAX_NR_EXERCISES]
	}
	out := []Exercise{}
	for _, ex := range e.Exercises {
		if strings.Contains(strings.ToLower(ex.Name), strings.ToLower(filter)) {
			out = append(out, ex)
			if len(out) >= MAX_NR_EXERCISES {
				break
			}
		}
	}

	return out
}

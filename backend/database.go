package backend

import (
	"embed"
	"encoding/json"
	"strings"
)

// This file contains the definition of the ExerciseDatabase struct and its methods.
// The ExerciseDatabase struct is used to store and retrieve exercises.
// It is currently just an embedded JSON file, but seems to work well enough.

type ExerciseDatabase struct {
	Exercises []Exercise
}

//go:embed exercises.json
var db_file embed.FS

// Initialize loads the exercises from the embedded JSON file.
func (e *ExerciseDatabase) Initialize() {
	data, _ := db_file.ReadFile("exercises.json")

	json.Unmarshal(data, &e.Exercises)
}

// getExerciseById returns the exercise with the given ID.
// If no exercise is found, it returns an empty Exercise struct.
func (e ExerciseDatabase) getExerciseById(id string) Exercise {
	for _, ex := range e.Exercises {
		if ex.Id == id {
			return ex
		}
	}
	return Exercise{}
}

// GetExercisesByIDs returns a list of exercises based on the given IDs.
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

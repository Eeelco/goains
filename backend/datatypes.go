package backend

// This file contains the definition of various datatypes used in the application.

// Exercise represents an exercise with its details.
type Exercise struct {
	Id               string   // Unique identifier for the exercise
	Name             string   // Name of the exercise
	Force            string   // Type of force applied during the exercise
	Level            string   // Difficulty level of the exercise
	Mechanic         string   // Movement mechanic of the exercise
	Category         string   // Category of the exercise
	Equipment        string   // Equipment required for the exercise
	PrimaryMuscles   []string // Muscles targeted as primary during the exercise
	SecondaryMuscles []string // Muscles targeted as secondary during the exercise
	Instructions     []string // Instructions for performing the exercise
	Images           []string // Images related to the exercise
}

// Set represents a set of repetitions and weight for an exercise.
type Set struct {
	Repetitions int     // Number of repetitions in the set
	Weight      float64 // Weight used for the set
}

// ExerciseUnit represents an exercise unit with its details.
type ExerciseUnit struct {
	ExerciseId   string // Identifier of the exercise
	ExerciseName string // Name of the exercise
	Rest         int    // Rest time in seconds after performing the exercise
	Sets         []Set  // Sets of repetitions and weight for the exercise
}

// PlanDay represents a day in a workout plan with its exercise units.
type PlanDay struct {
	Name          string         // Name of the day
	ExerciseUnits []ExerciseUnit // Exercise units for the day
}

// Plan represents a workout plan with its details and days.
type Plan struct {
	Name        string    // Name of the plan
	Description string    // Description of the plan
	Days        []PlanDay // Days in the plan
}

// Progress represents the progress of a workout plan.
type Progress struct {
	Start    string  // Start time of the workout
	Duration float64 // Duration of the workout in seconds
	Data     PlanDay // Data of the workout
}

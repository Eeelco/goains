package main

type Exercise struct {
	Id               string
	Name             string
	Force            string
	Level            string
	Mechanic         string
	Category         string
	Equipment        string
	PrimaryMuscles   []string
	SecondaryMuscles []string
	Instructions     []string
	Images           []string
}

type Set struct {
	Repetitions int
	Weight      int
}

type ExerciseUnit struct {
	ExerciseId   string
	ExerciseName string
	Rest         int
	Sets         []Set
}

type PlanDay struct {
	Name          string
	ExerciseUnits []ExerciseUnit
}

type Plan struct {
	Name        string
	Description string
	Days        []PlanDay
}

type Config struct {
	CurrentPlan   string
	NextDayIdx    int
	DefaultNrSets int
	DefaultNrReps int
	DefaultRest   int
}

func DefaultCfg() Config {
	return Config{
		CurrentPlan:   "",
		NextDayIdx:    0,
		DefaultNrSets: 3,
		DefaultNrReps: 10,
		DefaultRest:   60,
	}
}

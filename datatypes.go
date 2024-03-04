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
	Repetitions      int
	SecondsRestAfter int
}

type ExerciseUnit struct {
	ExerciseId string
	Sets       []Set
}

type PlanDay struct {
	Name          string
	ExerciseUnits []ExerciseUnit
}

type Plan struct {
	Name     string
	Days     []PlanDay
	DayOrder []int
}

type Config struct {
	CurrentPlan   string
	NextDayIdx    int
	DefaultNrSets int
	DefaultNrReps int
	DefaultPause  int
}

func DefaultCfg() Config {
	return Config{
		CurrentPlan:   "",
		NextDayIdx:    0,
		DefaultNrSets: 3,
		DefaultNrReps: 10,
		DefaultPause:  60,
	}
}

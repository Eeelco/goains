package main

const APP_NAME = "goainz"

var HOME_DIR string = ""
var CONFIG_DIR string = ""
var CONFIG_FILE string = ""

const exercises_json string = "https://raw.githubusercontent.com/yuhonas/free-exercise-db/main/dist/exercises.json"
const image_url_prefix string = "https://raw.githubusercontent.com/yuhonas/free-exercise-db/main/exercises/"

var exerciseDatabase []Exercise

package main

const APP_NAME = "goainz"
const MAX_NR_EXERCISES = 25

var HOME_DIR string = ""
var CONFIG_DIR string = ""
var CONFIG_FILE string = ""

var exerciseDatabase []Exercise

var config Config

# Goains

## About

A simple workout app built using Wails. Intended to be used on Linux mobile.
This project is very unfinished and under semi-active development.

## Configuration
Basic configuration can be done via the GUI. The config file and any created
workout plans are stored as json files in the `~/.config/goains` folder.

## Implemented features

- Create a workout plan, save it as a json file
- Included example workout
- Start a workout of the current plan

## Feature roadmap

- Better Readme (Screenshots, ...)
- Show which workout is next
- During workout (**Core functionality**):

  - Save progress afterwards

- Show more information about exercises (Description, images...)
- Track workout progress
- Ability to add own exercises to database
- UI cleanups
- Github action to automatically compile tagged releases
- (Maybe) Rename project

## Building

- Install and setup [Wails](https://wails.io)
- Run `wails build`
- To cross-compile for aarch64, run `make cross` (requires Docker)

## Credit

- [Wails](https://wails.io)
- [Free exercise DB](https://github.com/yuhonas/free-exercise-db)

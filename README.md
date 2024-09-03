# Goains

## About

A simple workout app built using Wails. Intended to be used on Linux mobile.
This project is very unfinished and under semi-active development.

## Configuration
Basic configuration can be done via the GUI. The config file and any created
workout plans are stored as json files in the `~/.config/goains` folder.

## Installation

### Download

Attached to each tagged release is a zip file containing compiled binaries for both aarch64
and amd64 architecture.

### Building from source

- Install and setup [Wails](https://wails.io)
- Run `wails build`
- To cross-compile for aarch64, run `make cross` (requires Docker)

## Implemented features

- Create a workout plan, save it as a json file
- Included example workout
- Start a workout of the current plan

## Feature roadmap

- Better Readme (Screenshots, ...)
- Notification sound when rest ends
- During workout (**Core functionality**):

  - Save progress afterwards

- Show more information about exercises (Description, images...)
- Track workout progress
- Ability to add own exercises to database
- UI cleanups
- (Maybe) Rename project

## Credit

- [Wails](https://wails.io)
- [Free exercise DB](https://github.com/yuhonas/free-exercise-db)

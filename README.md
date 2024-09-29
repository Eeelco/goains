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
- Save progress after workout

## Snapshots

![start](https://github.com/user-attachments/assets/bfb7a7db-e971-46ca-af2d-68266e0ea4f4)
![workout](https://github.com/user-attachments/assets/073aa7ae-f28c-471d-b9a8-c754bf4bee03)
![creator](https://github.com/user-attachments/assets/63856f07-53a8-417b-b0cb-ca43eb065d12)

## Feature roadmap

- Better Readme
- Notification sound when rest ends
- Show more information about exercises (Description, images...)
- Track workout progress
- Ability to add own exercises to database
- UI cleanups
- (Maybe) Rename project

## Credit

- [Wails](https://wails.io)
- [Free exercise DB](https://github.com/yuhonas/free-exercise-db)

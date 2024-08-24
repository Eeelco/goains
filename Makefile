pwd = $$(pwd)

dev: FORCE
	wails dev

build: FORCE
	wails build

cross: FORCE
	docker build -t goains_builder .
	docker run --rm -v $(pwd)/build/bin:/artifacts goains_builder

FORCE:
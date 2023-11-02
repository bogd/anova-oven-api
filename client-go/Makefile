.DEFAULT_GOAL := build
.PHONY: fmt lint vet mod build

fmt:
	cd src && go fmt ./...

lint: fmt
	cd src && golint ./...

vet: lint
	cd src && go vet ./...

mod:
	./update_mod.sh

compile:
	cd src && go build ./...

build: 
	./build_app.sh
	
run: build
	./binaries/anova-oven-api
	
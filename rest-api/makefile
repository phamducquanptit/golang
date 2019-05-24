SHELL := /bin/bash
BASEDIR = $(shell pwd)

run:
	swag init & go run main.go
doc:
	swag init
help: 
	@echo "make run -run file main.go"
	@echo "make docs -generate docs files"
# Makefile for site generation

.PHONY: all clean build

all: build

clean:
	@echo "Cleaning site/ output..."
	@rm -rf site

build: clean
	@echo "Generating site..."
	@go run ./cmd/sitegen

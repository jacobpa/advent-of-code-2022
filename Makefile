BUILD_DIRECTORY=build
BUILD=go build -v
BUILD_RELEASE=go build -v -ldflags "-s -w" -v
NIX_NAME=aoc2022
NAME=$(NIX_NAME)

ifeq ($(OS),Windows_NT)
	NAME := $(NIX_NAME).exe
endif

all: $(BUILD_DIRECTORY)/$(NAME)

$(BUILD_DIRECTORY)/$(NAME): main.go
	$(BUILD) -o $@

release_all: clean release_windows release_darwin release_linux

release_windows:
	GOOS=windows GOARCH=amd64 $(BUILD_RELEASE) -o $(BUILD_DIRECTORY)/windows/$(NIX_NAME).exe

release_darwin:
	GOOS=darwin GOARCH=arm64 $(BUILD_RELEASE) -o $(BUILD_DIRECTORY)/darwin/$(NIX_NAME)

release_linux:
	GOOS=linux GOARCH=amd64 $(BUILD_RELEASE) -o $(BUILD_DIRECTORY)/linux/$(NIX_NAME)

clean:
	rm -rf $(BUILD_DIRECTORY)/*

.PHONY: clean release_windows release_linux release_darwin

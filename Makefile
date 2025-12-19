.PHONY: all build run clean
.SILENT:

all: build

build:
	echo ***building***
	go build -o ./bin/ ./cmd/catfactsbot/

run: build
	echo ***runing***
	./bin/catfactsbot

clean:
	echo ***cleaning***
	rm ./bin/*

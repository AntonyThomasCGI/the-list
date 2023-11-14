.PHONY: all clean

# Variables

out := dist/the-list

all: build run

build:
	go build -o $(out) main.go

run:
	$(out)


clean:
	$(RM) -r dist/


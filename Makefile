.PHONY=build

build:
	go build ./cmd/jokes/

.DEFAULT:=build
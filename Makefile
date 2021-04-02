.PHONY: build examples clean

build:
	go build

examples:
	mkdir -p ./bin/
	go build -o ./bin/example ./examples/main.go

clean:
	rm -rf ./bin/

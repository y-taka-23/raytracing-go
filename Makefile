.PHONY: build clean

build:
	mkdir -p ./bin/
	go build -o ./bin/raytracing cmd/raytracing/main.go

clean:
	rm -rf ./bin/

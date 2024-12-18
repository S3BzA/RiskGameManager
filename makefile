all:
	go build -o bin/rgm *.go
	./bin/rgm

clean:
	rm -rf bin/*

run:
	./bin/rgm

build:
	go build -o bin/rgm *.go
all:
	go build -o bin/rgm main.go
	./bin/rgm

clean:
	rm -rf bin/*

run:
	./bin/rgm

build:
	go build -o bin/rgm main.go
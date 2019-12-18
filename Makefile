compile:
	go build .

run: compile
	./golang-kata-1	

test: compile
	go test -v ./...

build: test

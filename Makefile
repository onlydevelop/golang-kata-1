compile:
	go build .

run: compile
	./golang-kata-1	

test: compile
	go test -v ./...

cover: compile
	go test -v -cover ./...

report: compile
	go test -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html

build: test

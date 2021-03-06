all:	build

clean:
	rm -f mlfmt coverage.txt
	
format:
	gofmt -w .

test:
#	golint -set_exit_status ./...
	go vet ./...
#	errcheck ./...
#	go test ./... -v -covermode=atomic

build: clean format test
	go build

.PHONY: clean format test build

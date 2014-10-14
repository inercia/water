
# current version
VERSION=0.1.0

#################################################################
# main

all: dump

dump: FORCE
	@echo "Building water example app (dump)"
	go build github.com/inercia/water/cmd/dump

test: dump
	go test ./...

clean:
	@echo "Cleaning"
	@go clean
	rm -f dump
	rm -f *~ */*~

#################################################################
# deps

get: deps
deps:
	@echo "Getting all dependencies..."
	go get -d ./...

FORCE:

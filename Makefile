BINARY=gotracer

all:
	go build -gcflags=-trimpath="$(pwd)"

test:
	go test -v -gcflags=-trimpath="$(pwd)"

install:
	sudo cp -v "${BINARY}" /usr/local/bin/.

clean:
	go clean

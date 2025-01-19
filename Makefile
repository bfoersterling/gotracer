BINARY=gotracer

all:
	go build

test:
	go test -v

install:
	sudo cp -v "${BINARY}" /usr/local/bin/.

clean:
	go clean

BINARY=gotracer
GCFLAGS=-trimpath="$(shell pwd)"

all:
	go build -gcflags=${GCFLAGS}

test:
	go test -v -gcflags=${GCFLAGS}

install:
	sudo cp -v "${BINARY}" /usr/local/bin/.

clean:
	go clean

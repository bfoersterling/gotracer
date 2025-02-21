BINARY=gotracer
INSTALL_DIR=/usr/local/bin
GCFLAGS=-trimpath="$(shell pwd)"

all:
	go build -gcflags=${GCFLAGS}

test:
	go test -v -gcflags=${GCFLAGS}

install:
	sudo cp -v "${BINARY}" "${INSTALL_DIR}"

install_latest_release:
	wget https://github.com/bfoersterling/gotracer/releases/latest/download/gotracer_linux_amd64 -O /tmp/${BINARY}
	sudo chmod +x /tmp/${BINARY}
	sudo mv -v /tmp/${BINARY} /usr/local/bin/.

clean:
	go clean

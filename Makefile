default: build

build: build-server
	go build ./...

dev: build
	go-reload ./bookie-server

build-server:
	go build github.com/jgroeneveld/bookie2/cmd/bkserver
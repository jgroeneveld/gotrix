.PHONY: build build-server clean dev

build: web/views/compiled_ego.go build-server
	go build ./...

build-server:
	go build github.com/jgroeneveld/bookie2/cmd/bkserver

web/views/compiled_ego.go: web/views/*.ego
	go run cmd/ego/main.go -package=views -o $@ web/views
	goimports -w $@

clean:
	rm -f web/views/compiled_ego.go
	rm -f bkserver

dev: build
	go-reload ./bkserver
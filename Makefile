.PHONY: build build-server clean dev

build: web/frontend/views/compiled_ego.go build-server
	go build ./...

build-server:
	go build github.com/jgroeneveld/gotrix/cmd/gtserver

web/frontend/views/compiled_ego.go: web/frontend/views/*.ego
	go run cmd/ego/main.go -package=views -o $@ web/frontend/views
	goimports -w $@

clean:
	rm -f web/frontend/views/compiled_ego.go
	rm -f gtserver

dev: build
	go-reload ./gtserver
.PHONY: build build-server clean dev

build: app/web/frontend/views/compiled_ego.go build-server
	go build ./...

build-server:
	go build github.com/jgroeneveld/bookie2/cmd/bkserver

app/web/frontend/views/compiled_ego.go: app/web/frontend/views/*.ego
	go run cmd/ego/main.go -package=views -o $@ app/web/frontend/views
	goimports -w $@

clean:
	rm -f app/web/frontend/views/compiled_ego.go
	rm -f bkserver

dev: build
	go-reload ./bkserver
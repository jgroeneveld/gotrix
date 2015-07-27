.PHONY: build clean dev

build: web/frontend/views/compiled_ego.go
	go get ./...

web/frontend/views/compiled_ego.go: web/frontend/views/*.ego
	go run cmd/ego/main.go -package=views -o $@ web/frontend/views
	goimports -w $@

clean:
	rm -f web/frontend/views/compiled_ego.go

dev: build
	go-reload ./gtserver
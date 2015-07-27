.PHONY: build clean dev

build: web/frontend/views/compiled_ego.go web/frontend/assets/assets.go
	go get ./...

clean:
	rm -f web/frontend/views/compiled_ego.go
	rm -f web/frontend/assets/assets.go

dev: build
	go-reload gtserver

web/frontend/views/compiled_ego.go: web/frontend/views/*.ego
	go run cmd/ego/main.go -package=views -o $@ web/frontend/views
	goimports -w $@

ASSETS := $(shell find web/frontend/assets -type f | grep -v "^web/frontend/assets/assets.go$$")

web/frontend/assets/assets.go: $(ASSETS)
	@rm -f $@
	@go run cmd/goassets/main.go -file $@ web/frontend/assets

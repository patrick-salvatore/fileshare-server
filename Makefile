SHELL := /bin/bash

.PHONY: local
local:
	go build -o service cmd/main.go
	./service 

.PHONY: build
build:
	go build -o service cmd/main.go	

.PHONY: fresh
fresh:
	ENV=LOCAL \
	PORT=3001 \
	VERSION=VERSION \
	fresh
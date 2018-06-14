all: test build

run: build start

build:
	@echo " >> building binaries"
	@go build -o bin/trek-mp src/cmd/main-app.go

start:
	@echo " >> starting binaries"
	@./bin/trek-mp

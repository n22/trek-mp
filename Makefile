all: test build

run: build start

build:
	@echo " >> building binaries"
	@go build -o bin/trek-mp src/cmd/main-app.go

start:
	@echo " >> starting binaries"
	@./bin/trek-mp

pre-deploy:
	sudo mv bin/trek-mp /usr/local/bin/.
	sudo cp -r files/etc/trek-mp/. /etc/trek-mp/.
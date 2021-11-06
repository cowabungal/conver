.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

clear:
	rm -rf ./.bin

build-image:
	docker build -t conver .

start-container:
	docker run --env-file .env conver
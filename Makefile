APP_NAME=crypto-rate-notifier

.PHONY: run build docker-build docker-run

run:
	go run ./cmd/server

build:
	go build -o bin/$(APP_NAME) ./cmd/server

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p 8080:8080 $(APP_NAME)

clean:
	rm -rf bin/

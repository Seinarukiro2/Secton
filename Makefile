.PHONY: run
run:
	go run backend/main.go

.PHONY: migrate
migrate:
	go run backend/main.go migrate

.PHONY: db-up
db-up:
	docker-compose up -d

.PHONY: db-down
db-down:
	docker-compose down

.PHONY: build
build:
	go build -o bin/app backend/main.go

.PHONY: test
test:
	gotestsum --format=short-verbose -- ./backend/tests/...


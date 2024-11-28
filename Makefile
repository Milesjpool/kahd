default:
	@echo "Options:\n"\
		"build\n"\
		"start\n"\
		"start-local\n"\
		"test-unit\n"\
		"test-e2e"

build:
	docker build -t kahd-api-server --target api-server .

start-local:
	go run ./cmd/api-server

start: build
	docker run -p 8080:8080 api-server

test-unit:
	go test ./...

test-e2e: build
	docker build -t kahd-e2e-tests --target e2e-tests .
	docker-compose -f docker-compose.e2e.yml up --abort-on-container-exit
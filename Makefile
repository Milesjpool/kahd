default:
	@echo "Options:\n"\
		"build\n"\
		"start-local\n"\
		"test-e2e"

build:
	go build ./cmd/api-server

start-local:
	go run ./cmd/api-server

test-unit:
	go test ./...

test-e2e:
	go run e2e-tests

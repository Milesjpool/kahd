api = kahd-api
e2e-tests = kahd-e2e-tests

default:
	@echo "Options:\n"\
		"build-api\n"\
		"start-local\n"\
		"test-e2e"

build-api:
	cd $(api) \
	&& ./gradlew build \
	&& docker build -t $(api) .

start-local: build-api
	docker run -it -p 8080:8080 --rm $(api)

start-local-detach:
	$(eval api-container := $(shell docker run -d -p 8080:8080 --rm $(api)))
	@echo $(api-container)

test-e2e: build-api start-local-detach
	cd $(e2e-tests) \
	&& ./gradlew test -Denvironment=local -Dapi_version=$(shell cat $(api)/VERSION) \
	|| docker kill $(api-container)

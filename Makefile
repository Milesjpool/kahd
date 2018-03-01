api = kahd-api

default:
	@echo "Options:\n"\
		"build-api\n"\
		"start-api"\
	
build-api:
	cd $(api) \
	&& ./gradlew build \
	&& docker build -t $(api) .

start-api: build-api
	cd $(api) \
	&& docker run -it -p 8080:8080 --rm $(api)

api = kahd-api

api-image = eu.gcr.io/kahd-001/$(api)

default:
	@echo "Options:\n"\
		"build-api\n"\
		"start-api\n"\
		"push-api"

build-api:
	cd $(api) \
	&& ./gradlew build \
	&& docker build -t $(api) .

start-api: build-api
	cd $(api) \
	&& docker run -it -p 8080:8080 --rm $(api)

push-api: build-api
	cd $(api) \
	&& docker tag $(api) $(api-image) \
	&& gcloud docker -- push $(api-image)


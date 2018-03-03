api = kahd-api
version := $(shell cat $(api)/VERSION)
gcloud-project = kahd-001
api-image = eu.gcr.io/$(gcloud-project)/$(api):$(version)

default:
	@echo "Options:\n"\
		"build-api\n"\
		"start-api\n"\
		"push-api\n"\
		"deploy-api"

build-api:
	cd $(api) \
	&& ./gradlew build \
	&& docker build -t $(api) .

start-api: build-api
	docker run -it -p 8080:8080 --rm $(api)

push-api: build-api
	docker tag $(api) $(api-image) \
	&& gcloud docker -- push $(api-image)

deploy-api: push-api
	kubectl set image deployment/$(api) $(api)=$(api-image)

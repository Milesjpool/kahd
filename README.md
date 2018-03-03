# kahd
Card games API

## Requirements

* [Google Cloud SDK]
  *  Kubektl: `gcloud components install kubectl`
* [Docker]

## Setup

1. Clone this repository
1. Install dependencies
1. Login to Google Cloud CLI `gcloud auth login`
1. In the kahd-api directory, run `./gradlew`

## Building the project

1. From the repository root, run `make build-api`

## Pushing the project

1. From the repository root, run `make push-api`

[Google Cloud SDK]: https://cloud.google.com/sdk/docs/
[Docker]: https://docs.docker.com/install/

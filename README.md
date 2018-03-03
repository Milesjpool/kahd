[![Build Status](https://travis-ci.org/Milesjpool/kahd.svg?branch=master)](https://travis-ci.org/Milesjpool/kahd)

# kahd
Card games API

## Requirements

* [Git]
* [Google Cloud SDK]
  *  Kubektl: `gcloud components install kubectl`
* [Docker]

## Setup

1. Clone this repository
1. Install dependencies
1. Login to Google Cloud CLI `gcloud auth login`
1. In the kahd-api directory, run `./gradlew`

## Building & Deploying the API

1. From the repository root, run `make build-api`
1. Increment the API version in `kahd-api/VERSION`
1. From the repository root, run `make deploy-api`

[Git]: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
[Google Cloud SDK]: https://cloud.google.com/sdk/docs/
[Docker]: https://docs.docker.com/install/

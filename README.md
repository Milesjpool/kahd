[![Build and test](https://github.com/Milesjpool/kahd/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/Milesjpool/kahd/actions/workflows/build-and-test.yml)

# kahd
Open-source card games web-service including:
* API-server (WIP)
* GoLang API client

## Requirements
* [Git]
* [Google Cloud SDK]
  *  Kubectl: `gcloud components install kubectl`
* [Docker]

## Setup
1. Clone this repository
1. Login to Google Cloud CLI `gcloud auth login`
1. Run `./scripts/setup-cluster.sh` to set-up a gcloud cluster

## Building & Running the API locally
1. From the repository root, run `make start-local`

## Deploying the API
1. Increment the API version in `kahd-api/VERSION` (if required)
1. Push changes to master and [Travis][TravisCI build] will take it from there.

## Github Actions
This repository include a `.gihub/workflows` directory for testing and deploy the component(s).

## Useful resources
* [Setting up a containerised web-app in GCloud][Containerised web-app tutorial]

[Git]: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
[Google Cloud SDK]: https://cloud.google.com/sdk/docs/
[Docker]: https://docs.docker.com/install/

[Containerised web-app tutorial]: https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app

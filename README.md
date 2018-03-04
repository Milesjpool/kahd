[![Build Status](https://travis-ci.org/Milesjpool/kahd.svg?branch=master)](https://travis-ci.org/Milesjpool/kahd)

# kahd
Open-source card games web-service including:
* Kahd-API (WIP)

## Requirements

* [Git]
* [Google Cloud SDK]
  *  Kubectl: `gcloud components install kubectl`
* [Docker]

## Setup

1. Clone this repository
1. Install dependencies
1. Login to Google Cloud CLI `gcloud auth login`
1. In the kahd-api directory, run `./gradlew`
1. Run `./kahd-api/setup-cluster.sh` to set-up a gcloud cluster

## Building & Deploying the API

1. From the repository root, run `make build-api`
1. Increment the API version in `kahd-api/VERSION`
1. From the repository root, run `make deploy-api`

## TravisCI
This repository include a `.travis.yml` file to test and deploy the component(s).
Also included is a Travis-encrypted `credentials.tar.gz.enc` file ([described in this tutorial][GCloud Travis tutorial]). This contains:
* `client-secret.json` (for service-account access to GCloud).

## Useful resources
* [Setting up a containerised web-app in GCloud][Containerised web-app tutorial]
* [Integrating TravisCI with GCloud][GCloud Travis Tutorial]

[Git]: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
[Google Cloud SDK]: https://cloud.google.com/sdk/docs/
[Docker]: https://docs.docker.com/install/

[GCloud Travis tutorial]: https://cloud.google.com/solutions/continuous-delivery-with-travis-ci
[Containerised web-app tutorial]: https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app

![image](https://github.com/user-attachments/assets/72d61069-f0cf-443d-940d-1c5bdeff490f)


[![Build and test](https://github.com/Milesjpool/kahd/actions/workflows/build-and-test.yml/badge.svg?branch=main)](https://github.com/Milesjpool/kahd/actions/workflows/build-and-test.yml)
[![Smoke test](https://github.com/Milesjpool/kahd/actions/workflows/smoke-test.yml/badge.svg)](https://github.com/Milesjpool/kahd/actions/workflows/smoke-test.yml)

# kahd
Open-source card games web-service including:
* [API web-server][Kahd API] (WIP)

## Requirements
* [Git]
* [Google Cloud SDK]
  *  Kubectl: `gcloud components install kubectl`
* [Docker]

## Setup
1. Clone this repository
1. Login to Google Cloud CLI `gcloud auth login`
    1.  Access or create the relevant project.
1. Update `google_cloud/.env` as required for your project.
1. Run `./google_cloud/setup.sh` to set-up a gcloud stack

## Building & Running the API locally
1. From the repository root, run `make start-local`

## Deploying the API
1. Push changes to main and [GitHub Actions][GitHub Actions Build] will take it from there.

## Github Actions
This repository include a `.gihub/workflows` directory for testing and deploy the component(s).

## Useful resources
* [Setting up a containerised web-app in GCloud][Containerised web-app tutorial]

[Kahd API]: http://api.kahd.milesjpool.com
[Git]: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
[Google Cloud SDK]: https://cloud.google.com/sdk/docs/
[Docker]: https://docs.docker.com/install/
[GitHub Actions Build]: https://github.com/Milesjpool/kahd/actions/workflows/build-and-test.yml
[Containerised web-app tutorial]: https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app

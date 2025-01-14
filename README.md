# CI/CD Lecture template project

A simple checklist implemented in Golang with an angular frontend.

One day it will be a modern replacement for [strichliste.org](https://strichliste.org) maybe. 
Maybe check out their [demo](https://demo.strichliste.org) real quick to get a quick overview of the aimed feature set.

## How to create your own repo

```bash
# clone the original repo
git clone https://github.com/rubenhoenle/template.git

# create your github repo, set it to public, then get its (preferred ssh) clone url
git remote set-url origin git@github.com:yourusername/reponame.git
# then just push, that's it
git push
```

## Updating the repo after creation
```bash
# add second git remote ('origin' remote should be your GitHub repo)
git remote add upstream https://github.com/rubenhoenle/template.git

# fetch the latest changes from the new git remote
git fetch upstream

# check out your default branch
git checkout main

# rebase against the upstream "template" repository
git rebase upstream/main

# check everything, then push your changes
git push -f
```

## Usage

The app will run on port 8080. 

## Development

### Prerequisites

You should have installed the following tools on your system for development.

* [Golang](https://go.dev/doc/install) (version 1.23.4)
* [NodeJS](https://nodejs.org/en/download) (version 22 LTS)
* [Trivy](https://github.com/aquasecurity/trivy/releases)
* if possible: [Docker](https://docs.docker.com/get-started/get-docker/) (or if you know what you're doing you might also choose [Podman](https://podman.io/)).

For the users of the [nix package manager](https://nixos.org) among us, there is also a ´flake.nix´ provided which contains a devshell with all needed packages.

### Building the application

The backend provides a REST API which lets you create, update, delete, ... users, articles and transactions. 
The backend 

```bash
# building the frontend
cd frontend
npm ci
npx ng build
# copy the angular build result so it can be embedded by the go backend
cp -r frontend/dist/frontend/browser backend/cmd/strichliste/frontendDist

# running the application (e.g. for development)
cd backend
go run ./cmd/strichliste/main.go 
# you should be able to access the frontend of the app the app on http://localhost:8080 now in your browser
# the rest api should be accessible on http://localhost/api now

# building a binary for distribution
cd backend
CGO_ENABLED=0 go build -o ./strichliste ./cmd/strichliste/main.go
```

### Further information

If you want to get started with the REST API, please check out the [bruno](https://github.com/usebruno/bruno) collection which is part of the repository.

And here some further instructions which might be helpful for you during development:

```bash
# running all the backend tests
cd backend
go run ./...

# creating a test coverage report for the backend
cd backend
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html
```

### Code format

FYI: `gofmt` should already be available on your system if you have go installed.

```bash
# apply code format
gofmt -w ./backend
cd frontend; npx prettier --write .

# check code format
# FYI: this will print the list of files which are not formatted correctly
gofmt -l ./backend
cd frontend; npx prettier --check .

# check code format in CI pipeline
# FYI: the go fmt command does not end with exit code 1, instead prints only the list of files which are not formatted correctly. 
# So here's how to get a exit code 1 in case the format is not correct.
files=$(gofmt -l ./backend) && [ -z "$files" ]
```

### Smoke testing

To run the smoke tests, you will have to start an instance of the software first (native or via docker).

```bash
# run smoke tests on default url (http://localhost:8080)
cd backend
go run ./cmd/smoketest/

# run smoke test with url param
cd backend
go run ./cmd/smoketest/ --base-url http://localhost:9000
```

### Security scan

```bash
# run security scan with trivy, will produce report in html format (trivy-scan.html)
trivy fs --format template --output ./trivy-scan.html --template "@./trivy-html-report-template.tpl" .
```
# Pallpline

[![Pallpline](https://github.com/philipp-engelhardt/pallpline/actions/workflows/build.yml/badge.svg?branch=main&event=push)](https://github.com/philipp-engelhardt/pallpline/actions/workflows/build.yml)

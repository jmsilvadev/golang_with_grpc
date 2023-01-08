[![Build And Tests](https://github.com/jmsilvadev/golangtechtask/actions/workflows/tests.yml/badge.svg)](https://github.com/jmsilvadev/golangtechtask/actions/workflows/tests.yml)
[![Quality](https://github.com/jmsilvadev/golangtechtask/actions/workflows/quality.yml/badge.svg)](https://github.com/jmsilvadev/golangtechtask/actions/workflows/quality.yml)
[![Release](https://github.com/jmsilvadev/golangtechtask/actions/workflows/release.yml/badge.svg?branch=master)](https://github.com/jmsilvadev/golangtechtask/actions/workflows/release.yml)
![Min Coverage Badge](min-coverage.svg)
![Cur Coverage Badge](cur-coverage.svg) 


## Introduction

The main focus in the app design was to provide a clear solution with a consumer focus and a code with quality and estability, this means that it should provide a clear and simple solution to use. To achieve this, the SOLID (SRP, OCP, ISP and DIP) was followed, all the business logic was encapsulated in the pkg folder, having public visibility only of the methods that the consumer needs to use. Also were used the repository pattern and depency injection to grant the extensibility, reusability and facilitate the unit tests. To grant the quality was use tools of code style and automated tests with almost 100% of coverage and tests scenarios to prevent known and minimize unkown flaws.

## Tasks Done

- [x] Provide a `Go` implementation of the `GRPC` service in the `cmd/` directory of this repo.
- [x] Implement a `DynamoDB` based store for this `GRPC` service
- [x] Add pagination to the `ListVoteables` RPC call
- [x] Provide adequate test coverage for this simple service

1. [x] Adding Observability 

    Adding structured logging and/or tracing and metrics.
    (The current tech used should be considered when choosing technologies)
    - Used uber-go/zap to create structured logs and override the default log.

2. [x] Adding Configuration and Secrets management
    - Created a config package that handle the environments variables and injects in the server.

## CI/CD And SemVer

The project uses the Devops concept of continuous integration and continuous delivery through pipelines. To guarantee the CI there is a check inside the pipelines that checks the code quality and runs all the tests. This process generates an artifact that can be viewed and analyzed by developers.

To ensure the continuous delivery, the pipeline uses the automatic semantic versioning creating a versioned realease after each merge in the branch master. This release are tags in the control version system.

The semver system uses the angular commit message model [Angular Commit Message Format](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format).

## Usage

### Containerization

The project uses the concept of containerization which means that all the application is under dockers containers. To manipulate these containers please use the helpers showed in the next topic below.

### Helper

To facilitate the developers' work, a Makefile with containerized commands was created.

```bash
$ make
build                          Build docker image in daemon mode
cert                           Build a certificate to use with ssl
deps                           Install dependencies
doc                            Show package documentation
down                           Stop docker container
fmt                            Applies standard formatting
go-build                       Build a new binary server
lint                           Checks Code Style
logs                           Watch docker log files
proto                          Generate protobuffers
rebuild                        Rebuild docker container
ssh                            Interactive access to container
test.api                       Run end to end tests
test.coverage                  Run all available tests in a separate conteiner with coverage
test                           Run all available tests
tidy                           Dowload and Clean dependencies
up                             Start docker container
vendor                         Install vendor
vet                            Finds issues in code

```

### Configuration

The configuration of the application in under the file `docker-compose.yml`. Inside this file you can configure a serie of the environment variables like server port, aws credentials, etc.

### Build The Containers

To start the use of this application you need to build the images, to do this run the command:

```bash
$ make build

``` 

### Start The Server

To start the server run the command:
```bash
$ make up

``` 

After that, the server will be up and listen in the port configured in the config file.

### Tests

The project has unit tests, integration tests and end to end tests. To run these tests:
- Unit and Integration with coverage `make test.coverage`
- Unit and Integration `make test`
- End to End `make test.api`

#### Testing Manually

You can use the `grpcurl` to test manually the application. Follow some hints:

##### To create a Votable:

```bash
grpcurl -v --cert app/certs/server.crt --key app/certs/server.key --insecure -d '{"question":"What is the best football team in the universe?", "answers":["Flamengo","Arsenal","Chelsea", "Liverpool", "Manchester","Manchester City"]}' -proto app/api/service.proto localhost:4000 VotingService/CreateVoteable

Response contents:
{
  "UUID": "b5054e8b-ea55-480f-b272-0be1e7bdd6a9"
}

```

##### To cast a vote (use the response above)

```bash
grpcurl -v --cert app/certs/server.crt --key app/certs/server.key  --insecure -d '{"UUID":"b5054e8b-ea55-480f-b272-0be1e7bdd6a9","answer_index":0}' -proto app/api/service.proto localhost:4000 VotingService/CastVote

Response contents:
{
  "answer": "Flamengo",
  "answerVotes": "1"
}
```

Again :)

```bash

grpcurl -v --cert app/certs/server.crt --key app/certs/server.key  --insecure -d '{"UUID":"b5054e8b-ea55-480f-b272-0be1e7bdd6a9","answer_index":0}' -proto app/api/service.proto localhost:4000 VotingService/CastVote

Response contents:
{
  "answer": "Flamengo",
  "answerVotes": "2"
}

```

Ok, one more time :P

```bash
grpcurl -v --cert app/certs/server.crt --key app/certs/server.key  --insecure -d '{"UUID":"b5054e8b-ea55-480f-b272-0be1e7bdd6a9","answer_index":0}' -proto app/api/service.proto localhost:4000 VotingService/CastVote

Response contents:
{
  "answer": "Flamengo",
  "answerVotes": "3"
}

```

##### To list all votables

```bash
grpcurl -v --cert app/certs/server.crt --key app/certs/server.key  --insecure -proto app/api/service.proto localhost:4000 VotingService/ListVoteables

Response contents:
{
  "PreviousPage": "eyJVVUlEIjoiYjUwNTRlOGItZWE1NS00ODBmLWIyNzItMGJlMWU3YmRkNmE5In0=",
  "votables": [
    {
      "UUID": "b5054e8b-ea55-480f-b272-0be1e7bdd6a9",
      "question": "What is the best football team in the universe?",
      "answers": [
        "Flamengo",
        "Arsenal",
        "Chelsea",
        "Liverpool",
        "Manchester",
        "Manchester City"
      ],
      "votes": {
        "0": "3",
        "1": "0",
        "2": "0",
        "3": "0",
        "4": "0",
        "5": "0"
      }
    }
  ]
}
```

Question: Who is the best football team? :D

##### To list all votables with page size

Before we need to create more votables. Please return the creation section and create some votables.

Then, put a size as input to PageSize parameter (note: the order of votables is by creation desc):

```bash
grpcurl -v --cert app/certs/server.crt --key app/certs/server.key  --insecure -d '{"PageSize": 2}' -proto app/api/service.proto localhost:4000 VotingService/ListVoteables

Response contents:
{
  "NextPage": "eyJVVUlEIjoiNzQyNjZiYzgtNDUwNC00ODcwLWIxNmUtNjVmZmZjNjZlMjE5In0=",
  "PreviousPage": "eyJVVUlEIjoiMGE2MzFlNjYtY2Y4Ni00Njg5LTlhMDYtNDUzMDNlNDAwYTJmIn0=",
  "votables": [
    {
      "UUID": "0a631e66-cf86-4689-9a06-45303e400a2f",
      "question": "What is the best football team in the universe?",
      "answers": [
        "Flamengo",
        "Arsenal",
        "Chelsea",
        "Liverpool",
        "Manchester",
        "Manchester City"
      ],
      "votes": {
        "0": "0",
        "1": "0",
        "2": "0",
        "3": "0",
        "4": "0",
        "5": "0"
      }
    },
    {
      "UUID": "74266bc8-4504-4870-b16e-65fffc66e219",
      "question": "What is the best football team in the universe?",
      "answers": [
        "Flamengo",
        "Arsenal",
        "Chelsea",
        "Liverpool",
        "Manchester",
        "Manchester City"
      ],
      "votes": {
        "0": "0",
        "1": "0",
        "2": "0",
        "3": "0",
        "4": "0",
        "5": "0"
      }
    }
  ]
}

```

##### To list all votables with page size and pagination

Use the PreviousPage or NextPage tokens as input to Page parameter:

```bash
grpcurl -v --cert app/certs/server.crt --key app/certs/server.key  --insecure -d '{"PageSize": 2, "Page": "eyJVVUlEIjoiNzQyNjZiYzgtNDUwNC00ODcwLWIxNmUtNjVmZmZjNjZlMjE5In0="}' -proto app/api/service.proto localhost:4000 VotingService/ListVoteables

Response contents:
{
  "Page": "eyJVVUlEIjoiNzQyNjZiYzgtNDUwNC00ODcwLWIxNmUtNjVmZmZjNjZlMjE5In0=",
  "NextPage": "eyJVVUlEIjoiYjAyNTliYzMtNjg5Ni00OTBjLWIxM2YtNDk1NGExMWEwYjY0In0=",
  "PreviousPage": "eyJVVUlEIjoiMDRmZWVjNmUtYmUzYi00N2RlLWEwYmYtMmM5ZmRjZjlkZTZjIn0=",
  "votables": [
    {
      "UUID": "04feec6e-be3b-47de-a0bf-2c9fdcf9de6c",
      "question": "What is the best football team in the universe?",
      "answers": [
        "Flamengo",
        "Arsenal",
        "Chelsea",
        "Liverpool",
        "Manchester",
        "Manchester City"
      ],
      "votes": {
        "0": "0",
        "1": "0",
        "2": "0",
        "3": "0",
        "4": "0",
        "5": "0"
      }
    },
    {
      "UUID": "b0259bc3-6896-490c-b13f-4954a11a0b64",
      "question": "What is the best football team in the universe?",
      "answers": [
        "Flamengo",
        "Arsenal",
        "Chelsea",
        "Liverpool",
        "Manchester",
        "Manchester City"
      ],
      "votes": {
        "0": "0",
        "1": "0",
        "2": "0",
        "3": "0",
        "4": "0",
        "5": "0"
      }
    }
  ]
}

```

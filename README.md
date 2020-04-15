# crypto-currency-service
crypto-currency-service

## Development

### Prerequisites

* golang is *required* - version 1.13.x is tested; earlier versions may also
  work. See https://golang.org/doc/install
* docker is *required* - version 19.03.x is tested; earlier versions may also
  work.

###Build, Run service locally

```sh
navigate to crypto-currency-service
go build main.go 
go run main.go
service comes up with the port 8080 (default)

or
set env variables for CONFIG_REPO, ENV, PORT and CONFIG_REPO
go run main.go 
This will override default vaules
CONFIG_REPO is to use configuration file from remote host (github) etc..
```

###Running the tests

Run the tests locally with the following commands from the project root:

```sh
go test -v ./... -tags unit
go test -v ./... -tags integration
```

To build the image:
```sh
docker build --tag crypto-currency-service:1.0 .
```

To run docker container
```sh
 docker run docker run --publish 8000:8080 --detach --name crypto-currency-service crypto-currency-service:1.0
```

# Authentication Micro-service

[![Build Status](https://travis-ci.org/rij12/Authentication-Microservice.svg?branch=master)](https://travis-ci.org/rij12/Authentication-Microservice)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Auth_Service-GO&metric=alert_status)](https://sonarcloud.io/dashboard?id=Auth_Service-GO)
[![Go Report Card](https://goreportcard.com/badge/github.com/rij12/Authentication-Microservice)](https://goreportcard.com/report/github.com/rij12/Authentication-Microservice)
## Run

### Environment Vars 

```shell script

export MONGO_USERNAME=
export MONGO_PASSWORD=
export MONGO_HOST=
export MONGO_PORT=
export JWT_KEY=

```

## Setting up SSL 

Development Certs can be found in the Crypto folder. 
However, for production I would use Let's encrypt certs and store them in a secure location of the hosting server, in RAM. 


## Design 

## Docs

## TODO

* SSL 
* Tests...
* Refactor hard coded values into a config file. 
* Refactor Gobal Vars into a Singleton Config Service. 


# iot-demo
This repository demonstrates tools that can solve some fundamental problems with developing REST APIs in GoLang. Mostly using code generation.

## Toy Project
Will keep track of device metrics in a factory. Devices will periodically send metric data. According to the metric sent by the device, we may want to make the device sound an alert. We also want to be able to see metrics sent by a given device.

- Devices will be registered with serial number, registration date and a firmware version. They will be given a token on registration.
- Each device will send temperature data as a double value, using the token supplied on registration.
- When device sends a metric data, it may receive 'ALERT' or  'OK' message.
- We want to give an alert on a certain thresholds(initially between 70-100C). We want this configurable.
- We want an endpoint that accepts device id and returns metrics.

## Tools Demonstrated
- [wire](https://github.com/google/wire) for dependency injection with code generation
- [swag](https://github.com/swaggo/swag) for documentation generation
- [mockgen](https://github.com/golang/mock) for generating mock structs for interfaces 
- [viper](https://github.com/spf13/viper) for static/dynamic configuration. we use configmap mount with kubernetes in production. dynamic configuration works as is.


## Run the Server
Just run `go run iot-demo/cmd/server` to start the project. By default it listens to 8080 port. You can override it by setting the PORT env variable.

## Dependencies
To install dependencies used by the project, run `go mod download`. There are also some global dependencies that you may need to install to make `go generate ./...` run. Those are:

```
# for di
go get github.com/google/wire/cmd/wire
# for test mocks
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
# install swag for doc generation
go get github.com/swaggo/swag/cmd/swag
```

## Environment Variables
- `PROFILE` decides which environments configurations should override default configuration. Could be dev/stage/prod.
- `GO_ENV` can be production. Controls the `IsRelease` flag in global configuration.
- `PORT` decides which port to start the server on.

## Project Structure
```
.
├── adapters # implementation details
│   ├── config # static/dynamic config reading with viper
│   ├── http # REST HTTP API with gin, documented with swag
│   ├── jwt # jwt creation/parsing with jwt-go
│   └── storage # in memory repository functions (NOT THREAD SAFE!)
├── mocks # test struct mocks for interfaces
├── resources # configuration files
│   ├── config.yaml # static configuration file
│   └── config_dynamic.yaml # dynamic configuration file that will be re-read on changes
├── cmd
│   └── server # server entry point with wire
└─── pkg # business logic
    ├── auth
    ├── device
    │   ├── add-device # device registration use-case
    │   └── registry # device CRUD service
    └── metrics
        ├── query-metrics # metric query use-case
        ├── add-metrics # metric insert use-case
        ├── ingestion # metric CRUD service
        └── alert # threshold alert business logic

```

## Some Questions Related to the Structure
- Is it okay to put the mocks under `/mocks`. Any better place to put it?
- `adapters` folder includes implementation details of config/auth/http/repository. Should they be put under `/pkg` if so where? Putting them side by side with business logic, makes code harder to navigate.
- Better way of naming business logic related packages? Package under `/pkg/device/registry` is named `registry`. should it be named something else?


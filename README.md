# OMSHUB backend API (stub)

Go project using Gin router - https://github.com/gin-gonic/gin

## Setup

### VSCode fast-path

This project includes a [.devcontainers](https://code.visualstudio.com/docs/remote/containers) configuration
that can be used by VSCode to create a one-click development environment with Docker. The Docker container
includes all of the dependencies you need to compile Go, forwards the port exposed by HTTP server to your
local machine, and mounts the repository into the container so changes persist outside of Docker.

To get started:

1. Install the [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
   VSCode extension.
2. Open the repository with VSCode. You should see a prompt on the bottom left of the screen to open the
   project inside the container.

### Non-Docker/VSCode

1. Install [Go 1.18](https://go.dev/doc/install) and verify your installation by running `go version`.
2. Install [golangci-lint](https://golangci-lint.run/usage/install/#local-installation).
3. Install [Delve](https://github.com/go-delve/delve/tree/master/Documentation/installation).
2. Clone the repo.

## Development

### Run

1. Run `make build` to produce the executable binary.
2. Run `./core-api` to start the server.

### Test

1. Run `make test` to execute Go tests.

### Debug

1. Run `make debug` to debug the server. You can alternatively use the debugging capabilities provided by your editor.

## Directories

* `cmd/` - This is where the entrypoint into the application is (at `api/main.go`)
* `internal/` - This is where the routes of the server are defined. They are then imported into the main application in `cmd/api/main.go`

Currently there are only 2 routes defined, `/` and `/ping`.

## OpenAPI 3.x specification

* `doc/openapi.yaml` - This is a sample specification 3.0.3 doc, more examples can be found in https://github.com/OAI/OpenAPI-Specification
* The yaml file can be imported to [ReDoc](https://redocly.github.io/redoc/) and Postman
# core-api

https://api.omshub.org

This repository contains the backend of OMShub, dubbed the "Core API". It also contains (temporarily) Terraform
infrastructure-as-code definitions in the `ops/` subdirectory.

**NB: some of the following is forward-looking and may not reflect the current state of the system.**

The Core API is a Go service backed by a Postgres database, deployed onto DigitalOcean's [App Platform](https://docs.digitalocean.com/products/app-platform/)
PaaS. It exposes a REST API defined with a [Swagger/OpenAPI specification](https://swagger.io/specification/) which
the NextJS frontend queries. Users authenticate against the API using an OAuth 2.0 scheme, by supplying a bearer
JWT issued by Auth0.

The Go service uses the [Gin framework](https://github.com/gin-gonic/gin) for routing and HTTP handling. It uses
middleware to push basic metrics to NewRelic.

## Setup

### VSCode fast-path

This project includes a [.devcontainers](https://code.visualstudio.com/docs/remote/containers) configuration
that can be used by VSCode to create a one-click development environment with Docker. The Docker container
includes all of the dependencies you need to compile Go and start Postgres. It forwards the ports exposed
by the HTTP server and Postgres DB to your local machine, and mounts the repository into the container so
changes persist outside of Docker.

To get started:

1. Install the [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
   VSCode extension.
2. Open the repository with VSCode. You should see a prompt on the bottom left of the screen to open the
   project inside the container.

### Non-Docker/VSCode

1. Install [Go 1.18](https://go.dev/doc/install) and verify your installation by running `go version`.
2. Install [golangci-lint](https://golangci-lint.run/usage/install/#local-installation).
3. Install [Delve](https://github.com/go-delve/delve/tree/master/Documentation/installation).
4. Install Postgres. Set it up according to the configuration in `.devcontainer/.env`.

## Development

### How do I run the Go service?

1. Run `make build` to produce the executable binary.
2. Run `./core-api` to start the server.

### How do I test the Go service?

1. Run `make test` to execute Go tests.

### How do I debug the Go service?

1. Run `make debug` to debug the server. You can alternatively use the debugging capabilities provided by your editor.

## Deployment

Terraform will automatically provision any required infrastructure changes (this should only happen if something under
`ops/` is modified) when a PR is merged. Following a successful Terraform apply, the Go service will be deployed to
DigitalOcean.

Tip: the Terraform workflow will comment on your PR with an overview of any infrastructure changes that are needed.

## Additional information

- [ARCHITECTURE.md](./ARCHITECTURE.md).

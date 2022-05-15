FROM mcr.microsoft.com/vscode/devcontainers/go:1.18-bullseye

RUN apt update \
    && apt upgrade -y

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s 

# Delve for debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# goimports for source formatting
RUN go install golang.org/x/tools/cmd/goimports@latest

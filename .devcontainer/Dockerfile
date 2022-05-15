FROM mcr.microsoft.com/vscode/devcontainers/go:1.18-bullseye

RUN apt update \
    && apt upgrade -y

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s 
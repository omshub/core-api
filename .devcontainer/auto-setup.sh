#!/bin/bash

set -euo pipefail

# golangci-lint for debugging linting issues locally
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s 

# Delve for debugging
go install github.com/go-delve/delve/cmd/dlv@latest

# goimports for source formatting
go install golang.org/x/tools/cmd/goimports@latest

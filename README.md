# OMSHUB backend API (stub)

Go project using Gin router - https://github.com/gin-gonic/gin

## Setup
1. Install Go: https://go.dev/doc/install
2. Clone the repo
3. Run `go run cmd/api/main.go`


## Directories
* `cmd/` - This is where the entrypoint into the application is (at `api/main.go`)
* `internal/` - This is where the routes of the server are defined. They are then imported into the main application in `cmd/api/main.go`

Currently there are only 2 routes defined, `/` and `/ping`.


## Other things of interest
* This project uses Go modules for dependency management, take a look at the files `go.mod`, `go.sum`, and the directory `vendor/`. You can learn a bit more about Go modules here: [DigitalOcean Go Modules Introduction](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules), [Go Blog Official Go Modules Intro](https://go.dev/blog/using-go-modules).
* 
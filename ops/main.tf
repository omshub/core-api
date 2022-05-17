terraform {
  cloud {
    organization = "omshub"
    workspaces {
      name = "core-api"
    }
  }

  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

variable "do_token" {}

provider "digitalocean" {
  token = var.do_token
}

variable "newrelic_api_key" {}

resource "digitalocean_app" "app-core-api" {
  spec {
    name   = "core-api"
    region = "nyc1"

    service {
      name               = "core-api"
      instance_count     = 1
      instance_size_slug = "basic-xxs"

      http_port = 1927

      dockerfile_path = "Dockerfile"

      routes {
        path = "/"
      }

      env {
        key   = "NEWRELIC_API_KEY"
        value = var.newrelic_api_key
        scope = "RUN_TIME"
        type  = "SECRET"
      }

      env {
        key   = "NEWRELIC_APP_NAME"
        value = "omshub/core-api"
        scope = "RUN_TIME"
      }

      github {
        repo           = "omshub/core-api"
        branch         = "ci/digitalocean-app-platform"
        deploy_on_push = false
      }
    }
  }
}
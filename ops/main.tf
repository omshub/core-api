terraform {
  # We're using Terraform Cloud as the .tfstate backend and the
  # pipeline runner.
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

# These are set by GitHub Actions from repository secrets.
variable "do_token" {}
variable "newrelic_api_key" {}

provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_app" "app_core_api" {
  spec {
    name   = "core-api"
    region = "nyc"

    # CNAME record needs to be manually configured in Google Domains
    # to route to the do_app_url TF output var.
    domain {
      name = "api.omshub.org"
    }

    service {
      name               = "core-api"
      instance_count     = 1
      instance_size_slug = "basic-xxs"

      # DO sets the PORT env var to this value.
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

      # DO pulls from this repo to build and deploy.
      github {
        repo           = "omshub/core-api"
        branch         = "main"
        deploy_on_push = false
      }
    }
  }
}

output "do_app_url" {
  value = digitalocean_app.app_core_api.live_url
}

output "do_app_deployment_id" {
  value = digitalocean_app.app_core_api.active_deployment_id
}

# Resources for Terraform Cloud

resource "tfe_organization" "org-omshub" {
  name  = "omshub"
  email = "terraform@omshub.org"
}

resource "tfe_workspace" "workspace-core-api" {
  name         = "core-api"
  organization = tfe_organization.org-omshub.name
}

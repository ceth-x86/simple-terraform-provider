terraform {
  required_providers {
    simpleprovider = {
      version = "0.1"
      source  = "hashicorp.com/edu/simpleprovider"
    }
  }
}

provider "simpleprovider" {}

module "psl" {
  source = "./entity"
  entity_name = "Just sample entity name"
}

output "psl" {
  value = module.psl.entity
}

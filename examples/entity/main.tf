terraform {
  required_providers {
    simpleprovider = {
      version = "0.1"
      source  = "hashicorp.com/edu/simpleprovider"
    }
  }
}

variable "entity_name" {
  type    = string
  default = "Default entity name"
}

data "simpleprovider_entities" "all" {}

# Returns all coffees
output "all_entities" {
  value = data.simpleprovider_entities.all
}

# Only returns packer spiced latte
output "entity" {
  value =  data.simpleprovider_entities.all
}

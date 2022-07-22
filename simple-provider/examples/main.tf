terraform {
  required_providers {
    simpleprovider = {
      version = "0.1"
      source  = "hashicorp.com/edu/simpleprovider"
    }
  }
}

provider "simpleprovider" {}

resource "simpleprovider_item" "test" {
  name = "one"
  description = "first"
}
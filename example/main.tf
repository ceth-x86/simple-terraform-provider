terraform {
  required_providers {
    example = {
      version = "0.1"
      source  = "hashicorp.com/edu/example"
    }
  }
}

provider "example" {
}

resource "example_item" "test" {
  name = "this_is_an_item"
  description = "this is an item"
}

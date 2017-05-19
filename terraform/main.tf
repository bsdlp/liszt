provider "aws" {
  region = "us-west-2"
}

module "liszt" {
  source = "./modules"
  region = "us-west-2"
  env    = "testing"
}

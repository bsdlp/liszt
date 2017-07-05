provider "aws" {
  region = "us-west-2"
}

module "testing" {
  source = "./modules"
  region = "us-west-2"
  env    = "testing"
}

module "dev" {
  source = "./modules"
  region = "us-west-2"
  env     = "dev"
}

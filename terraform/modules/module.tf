variable "region" {
  default = "us-west-1"
}

variable "env" {
  default = "testing"
}

provider "aws" {
  region = "${var.region}"
}

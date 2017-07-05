variable "region" {
  default = "us-west-2"
}

variable "env" {
  default = ""
}

provider "aws" {
  region = "${var.region}"
}

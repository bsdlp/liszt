terraform {
  backend "s3" {
    bucket = "bsdlp"
    key    = "tfstate/liszt/terraform.tfstate"
    region = "us-west-2"
  }
}

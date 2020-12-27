terraform {
    backend "s3" {
        bucket         = "lists-terraform-state"
        key            = "global/s3/terraform.tfstate"
        region         = "eu-west-3"
    }
}

provider "aws" {
    profile = "default"
    region = "eu-west-3"
}

resource "aws_elastic_beanstalk_application" "tftest" {
  name        = "lists-${terraform.workspace}-001"
  description = "lists backend api"
}
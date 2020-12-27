terraform {
    backend "s3" {
        bucket         = "lists-terraform-state"
        key            = "global/s3/terraform.tfstate"
        region         = "eu-west-3"
    }
}

locals {
    appversion = "001"
    appname = "lists-backend-${terraform.workspace}-${local.appversion}"
}

provider "aws" {
    profile = "default"
    region = "eu-west-3"
}

resource "aws_elastic_beanstalk_application" "application" {
  name        = local.appname
}

resource "aws_elastic_beanstalk_environment" "environment" {
  name                = "${local.appname}-environment"
  application         = aws_elastic_beanstalk_application.application.name
  solution_stack_name = "64bit Amazon Linux 2 v3.1.3 running Go 1"

    setting {
        namespace = "aws:autoscaling:launchconfiguration"
        name      = "IamInstanceProfile"
        value     = "aws-elasticbeanstalk-ec2-role"
    }
}


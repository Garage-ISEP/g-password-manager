terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  backend "s3" {}


  required_version = ">= 1.2.0"
}

provider "aws" {
  region  = "eu-west-3"
}

locals {
  prefix = var.stack-name
  suffix = var.env
  tags = {
    project      = var.stack-name
    environement = var.env
  }
}



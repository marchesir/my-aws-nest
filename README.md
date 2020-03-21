# AWS Local Dev Environment
AWS local dev environment good practices.

## Introduction
This is a set of gneral good practices for setting up AWS for local development from private users point of view not commercial.  As we are dealing with pure Infra Terraform will be used to achieve this.

## Objective 
The final goal is to not use the AWS root account and have 2 dedicated accounts and VPC's for each one, this is general good practive from a security point of view.  The basic concept each account should only be able to do what is reuired (least pridledge model).

## Steps

Assumption: AWS had been setup with MFA as reccomended and Terraform setup locally ([Terraform Setup](https://learn.hashicorp.com/terraform/getting-started/install.html)).  Later all local dev will reside only on AWS, but for initial setup local dev pc is required.

1. Login into root account and create 2 accounts as such:
   * <name>dev: eg. richardev
   * <name>admin: eg richardadmin

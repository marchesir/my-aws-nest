## Introduction
Repo containing resoures for learning Kubernetes (K8s) from the bottom up.

## Objectives
Use https://github.com/kelseyhightower/kubernetes-the-hard-way as a starting point to lean the inner workings of K8s.

## Environment
AWS: simple setup:
1. root account with MFA: do not touch;
2. admin user with MFA: to be used to all work; *
3. dedicated restricted API only (no console access) for K8s resources;

* to be split into admin only account and dev account at some point in future.

## Steps

1. from admin accountutilize AWS cloud shell for work;
2. initally we will use bash to "automate infra";

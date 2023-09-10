# CI/CD Play

## Heads Up

Most of what I'm doing in this repo is offline, so there's a lot more that I haven't uploaded, including more intricate golang modules, GCP infra (instead of Linode), and a few other half implented tools. This playground is very much a scratch space...

## What this repo **is**

This is a play space for me to be able play with building a full devops pipeline from the ground up. At present I've implemented the following:

* [x] Terraform IaC
* [x] Kubernetes cluster on Linode
* [x] Pods distributed with load balancing service
* [x] GitHub actions workflows to containerise and push new images to docker
* [x] Python Test Cases
* [ ] Separate container within each pod communicates with main container
* [ ] Centralised DB service pod/volume
* [ ] Promethius logging

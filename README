# UME Devops Test

## Setup

The following tools a required to run this project:
* [Docker Engine](https://docs.docker.com/engine/install/ubuntu/)
* [Kind kubernetes local cluster](https://kind.sigs.k8s.io/docs/user/quick-start)
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)
* [helm](https://helm.sh/docs/intro/install/)
* [kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/)

## Deploy local

To create the local cluster you will need to run the following.
```bash
make setup-cluster-local
```

This will create the a Kind cluster with the name `devcluster` and it will install metallb, istio and cert-manager on the local cluster.

It will also create two Istio gateways `local.example.com.br` and `local.internal.example.com.br`. they represent a pubic gateway, 
exposed on the internet, and a internal gateway, available only on the private network, respectively.

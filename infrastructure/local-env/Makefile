SHELL := bash

setup-cluster-local: 
	envsubst < kind-config.yaml > kind-config-tmp.yaml
	kind create cluster --config kind-config-tmp.yaml --name devcluster && \
		kubectl config set-context kind-devcluster
	rm -rf kind-config-tmp.yaml
	kubectl config set-context kind-devcluster && cd metallb && $(MAKE) deploy
	kubectl config set-context kind-devcluster && cd ./../cert-manager && $(MAKE) deploy
	kubectl config set-context kind-devcluster && cd ./../istio && $(MAKE) deploy
	
destroy-cluster-local:
	 kind delete cluster --name devcluster

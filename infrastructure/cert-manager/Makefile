SHELL := bash

deploy:
	kubectl create namespace cert-manager
	cd operator/local && kustomize build --enable-helm | kubectl apply -f -
	kubectl wait --namespace cert-manager --for=condition=ready pod --selector=app.kubernetes.io/instance=cert-manager --timeout=90s
	cd cluster-issuer/local && kustomize build --enable-helm | kubectl apply -f -

destroy:
	cd operator/local && kustomize build --enable-helm | kubectl delete -f -
	cd cluster-issuer/local && kustomize build --enable-helm | kubectl delete -f -
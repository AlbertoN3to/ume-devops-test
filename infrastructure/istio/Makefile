SHELL := bash


deploy:
	kubectl create namespace istio-system
	kubectl create namespace istio-ingress
	cd base/local && kustomize build --enable-helm | kubectl apply -f -
	cd discovery/local && kustomize build --enable-helm | kubectl apply -f -
	cd gateway/local && kustomize build --enable-helm | kubectl apply -f -

destroy:
	cd gateway/local && kustomize build --enable-helm | kubectl delete -f -
	cd discovery/local && kustomize build --enable-helm | kubectl delete -f -
	cd base/local && kustomize build --enable-helm | kubectl delete -f -
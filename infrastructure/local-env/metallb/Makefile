deploy:
	kubectl config set-context kind-devcluster && \
		kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml > /dev/null 2>&1
	echo "Waiting for metal load balancer instalation..." && sleep 30
	kubectl config set-context kind-devcluster && kubectl wait --namespace metallb-system --for=condition=ready pod --selector=app=metallb --timeout=90s
	export PRE_FIX=$(shell docker network inspect -f '{{(index .IPAM.Config 0).Subnet}}' kind | cut -d '.' -f -2); \
		kubectl config set-context kind-devcluster && envsubst < ip-adress-pool.yaml | kubectl apply -f - > /dev/null 2>&1

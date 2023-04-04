## get docker kind network to add to the configmap
docker network inspect -f '{{(index .IPAM.Config 0).Subnet}}' kind | cut -d '.' -f -2

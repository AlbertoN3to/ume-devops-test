SHELL := bash

setup-cluster-local:
	cd infrastructure/local-env && $(MAKE) -B setup-cluster-local

destroy-cluster-local:
	cd infrastructure/local-env && $(MAKE) -B destroy-cluster-local

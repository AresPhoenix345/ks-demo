# ArgoCD Integration Makefile

.PHONY: build test generate

build:
	go build -o bin/argocd-integration-controller ./cmd/controller

test:
	go test ./...

generate:
	@echo "Run controller-gen if needed (e.g. make -C ../../hack generate)"
	@echo "Skipping for standalone prototype."

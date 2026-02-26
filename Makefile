# Senior Engineering Workflow Makefile

# Variables
NAMESPACE=dev
K8S_PATH=Academy/Week-18-Kubernetes/k8s/overlays/dev
LB_URL=http://localhost:8000

.PHONY: deploy check clean logs help

## deploy: Apply Kustomize manifests to the cluster
deploy:
	@echo "🚀 Deploying system to Kubernetes (Namespace: $(NAMESPACE))..."
	kubectl apply -k $(K8S_PATH)

## check: Wait for pods to be ready and verify system health
check:
	@echo "⏳ Waiting for backend pods to be ready..."
	kubectl wait --for=condition=ready pod -l app=backend -n $(NAMESPACE) --timeout=60s
	@echo "⏳ Waiting for load balancer pod to be ready..."
	kubectl wait --for=condition=ready pod -l app=lb -n $(NAMESPACE) --timeout=60s
	@echo "🔍 Testing Load Balancer response..."
	@curl -s "$(LB_URL)?user_id=makefile_check" || (echo "❌ Failed to connect to LB. Ensure port-forwarding is active if on Minikube." && exit 1)
	@echo "✅ SYSTEM ONLINE: Load Balancer and Backends are communicating successfully."

## logs: Tail logs from the Load Balancer
logs:
	kubectl logs -f deployment/lb-deployment -n $(NAMESPACE)

## clean: Tear down the infrastructure
clean:
	@echo "🧹 Removing Kubernetes resources..."
	kubectl delete -k $(K8S_PATH)

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^##' Makefile | sed -e 's/## //'

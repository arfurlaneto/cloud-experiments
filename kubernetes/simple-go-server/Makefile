CLUSTER_NAME=mycluster
IMAGE_NAME=arfurlaneto/goserver

docker-build:
	docker build --build-arg APP_VERSION=1.0 -t ${IMAGE_NAME}:1 .
	docker build --build-arg APP_VERSION=2.0 -t ${IMAGE_NAME}:2 .
	docker build --build-arg APP_VERSION=3.0 -t ${IMAGE_NAME}:3 -t ${IMAGE_NAME}:latest .

docker-run:
	docker run --rm -p 80:8000 ${IMAGE_NAME}:latest

docker-prune:
	docker image remove -f $$(docker image list -q --filter "reference=arfurlaneto/goserver") | true

kind-create:
	kind create cluster --config=kind/kind.yaml --name=${CLUSTER_NAME}
	kubectl config use-context kind-${CLUSTER_NAME}

kind-destroy:
	kind delete clusters ${CLUSTER_NAME}

kind-load-images: docker-build
	kind load docker-image ${IMAGE_NAME}:1 -n ${CLUSTER_NAME}
	kind load docker-image ${IMAGE_NAME}:2 -n ${CLUSTER_NAME}
	kind load docker-image ${IMAGE_NAME}:3 -n ${CLUSTER_NAME}
	kind load docker-image ${IMAGE_NAME}:latest -n ${CLUSTER_NAME}

k8s-setup:
	kubectl apply --filename https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
	kubectl wait --namespace ingress-nginx \
		--for=condition=ready pod \
		--selector=app.kubernetes.io/component=controller \
		--timeout=90s
	kubectl apply -f k8s

deploy:
	bash kind.sh
	kubectl label namespace default observe=true
	helm install --repo https://helm.traefik.io/traefik traefik traefik
	kubectl create namespace logger
	make logger
	make certs
	kubectl apply -k ./certs
	kubectl apply -f ./logger/logger.yaml
	kubectl wait --namespace=logger --timeout=5m --for=condition=Ready=True -l app=logger pod
	./webhook.sh
	@sleep 10
	kubectl apply -f deployment.yaml
	kubectl wait --timeout=5m --for=condition=Ready=True -l app.kubernetes.io/name=traefik pod
	kubectl port-forward svc/traefik 8080:80

modinfo:
	@docker run --name jobs traefik/jobs:helmsman > /dev/null || true
	@docker cp jobs:/start jobs > /dev/null
	@docker rm jobs > /dev/null
	@go run ./modextractor/main.go jobs
	@rm jobs

.PHONY: logger
logger:
	docker build -t localhost:5000/logger ./logger
	docker push localhost:5000/logger

.PHONY: certs
certs:
	mkcert -cert-file certs/cert.pem -key-file certs/key.pem logger.logger.svc

clean:
	kind delete cluster
	docker stop kind-registry
	docker rm kind-registry
	rm ./certs/cert.pem
	rm ./certs/key.pem

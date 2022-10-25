
.PHONY: run-server
run-server:
	export $(cat ./server/.env | xargs)
	go run ./server/main.go

.PHONY: docker-build-server
docker-build-server:
	docker build --rm -t demo-circuit-breaker-server . -f server/Dockerfile

.PHONY: run-client
run-client:
	export $(cat ./client/.env | xargs)
	go run ./client/main.go

.PHONY: docker-build-client
docker-build-client:
	docker build --rm -t demo-circuit-breaker-client . -f client/Dockerfile

.PHONY: run
run: docker-build-server docker-build-client
	docker compose up --force-recreate
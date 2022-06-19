.PHONY: run
run:
	docker compose up -d

.PHONY: stop
stop:
	docker compose stop

.PHONY: lint
lint:
	docker compose exec server golangci-lint run

.PHONY: test
test:
	docker compose exec server go test ./... -count=1

.PHONY: generate
generate:
	docker compose exec server go generate ./...

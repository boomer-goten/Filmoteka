.PHONY: run_service
run_service:
	@go run ./cmd/filmoteka-server/main.go --config=./config/config.yaml

.PHONY: compose
compose:
	@docker-compose up --build filmoteka-server

.PHONY: test
test:
	@go test
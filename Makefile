install-tools:
	@echo installing tools && \
	@go install \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@echo done

generate:
	@echo running code generation
	@go generate ./...
	@echo done

.PHONY: up
up: ## run the application on docker
	@docker-compose up --build -d

.PHONY: db
db: ## create the database
	@docker exec -it postgres createdb --username=root --owner=root event_driven

.PHONY: drop-db ## drop the database
	@docker exec -it postgres dropdb event_driven

.PHONY: migration
migration: ## create new migration file
	@migrate create -ext sql -dir db/migrations -seq $(name)

.PHONY: migrate
migrate: ## apply all up migrations
	@migrate -source file://db/migrations -database postgres://root:secret@localhost:5433/event_driven?sslmode=disable up $(version)

.PHONY: migrate-down
migrate-down: ## apply all down migrations
	@migrate -source file://db/migrations -database postgres://root:secret@localhost:5433/event_driven?sslmode=disable down $(version)

.PHONY: test
test: ## unit testing with coverage
	@go test -v -cover ./...
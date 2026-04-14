# Default recipe to list all available recipes
default:
    @just --list

# Database migrations
migrate:
    go run cmd/migrate/main.go up

migrate-down:
    go run cmd/migrate/main.go down

# Protobuf generation
generate-proto:
    protoc --proto_path=pkg/proto --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative pkg/proto/*.proto

# Code quality
fmt:
    go fmt ./...

vet:
    go vet ./...

lint:
    golangci-lint run

# Running the application
run-client:
    go run cmd/client/main.go

run-server:
    go run cmd/server/main.go


test-repo:
    go test -race -v ./tests/repository/...

test-service:
    go test -race -v ./tests/service/...

test-handler:
    go test -race -v ./tests/handler/...

test-integration: test-repo test-service test-handler

test-all:
    go test -race -covermode=atomic -coverprofile=coverage.txt -v ./...

# Coverage and Tools
coverage:
    go tool cover -html=coverage.txt

sqlc-generate:
    sqlc generate

db-docs:
    dbdocs build doc/db.dbml

db-schema:
    dbml2sql --postgres -o doc/schema.sql doc/db.dbml

db-sqltodbml:
    sql2dbml --postgres doc/schema.sql -o doc/db.dbml

generate-swagger:
    swag init -g cmd/client/main.go

# Docker
docker-up:
    docker compose up -d --build

docker-down:
    docker compose down

# k6 Performance Testing
k6-test domain script="common":
    k6 run k6/{{domain}}/{{domain}}_{{script}}.js

migrate:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

generate-proto:
	protoc --proto_path=pkg/proto --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative pkg/proto/*.proto

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

run-client:
	go run cmd/client/main.go

run-server:
	go run cmd/server/main.go

test:
	go test -race -covermode=atomic -coverprofile=coverage.txt -v ./tests/unit_test/...


test-all:
	go test -race -covermode=atomic -coverprofile=coverage.txt -v ./...

coverage:
	go tool cover -html=coverage.tx

sqlc-generate:
	sqlc generate

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

db_sqltodbml:
	sql2dbml --postgres doc/schema.sql -o doc/db.dbml

generate-swagger:
	swag init -g cmd/client/main.go

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down

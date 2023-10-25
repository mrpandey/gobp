#!/usr/bin/env just --justfile

# load variables in .env to environment
set dotenv-load

GO := "go"
GOVET_CMD := GO + " vet"
GOTEST_CMD := GO + " test"
GOCOVER_CMD := GO + " tool cover"
GOBUILD_CMD := GO + " build"
GORUN_CMD := GO + " run"

COV_THRESHOLD := "95"

DB_URI := "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB"
DEV_DB_URI := "postgres://postgres:postgres@localhost:5432/temp_migration_dev"
MIGRATION_DIR := "file://migrations?format=golang-migrate"
DB_URI_PARAMS := if env_var_or_default("POSTGRES_SSL_MODE", "prefer") == "disable" { "?search_path=public&sslmode=disable" } else { "?search_path=public" }


# display all commands
list:
    @just --list --unsorted

# run tests in verbose mode (otherwise panics go undetected)
test:
    just run-docker postgres
    {{GOTEST_CMD}} -v -failfast -race -covermode=atomic -coverprofile=coverage.out ./...
    @{{GOCOVER_CMD}} -func=coverage.out
    @{{GOCOVER_CMD}} -html=coverage.out -o coverage.html

# update go project dependencies
tidy:
    go get -u ./...
    go mod tidy -v

# format long lines in source
golines path:
    golines -m 120 -w {{path}}

# run static checks
lint:
    golangci-lint run --fix



# run service natively with hot-reloading. service: "server" | "worker"
run service:
    just runc postgres
    just proto-gen
    air -c .air.{{service}}.toml

# run multiple docker-compose services
runc +SERVICES:
    docker-compose -f dev/docker-compose.yml --project-name gobp up -d {{SERVICES}}

# build binary for the given service
build service:
    rm -f ./build/gobp_{{service}} && CGO_ENABLED=0 {{GOBUILD_CMD}} -ldflags="-w -s" -o ./build/gobp_{{service}} ./src/cmd/{{service}}

# build docker images for specified services; see docker-compose.yml for list of services
build-docker +SERVICES:
    docker-compose -f dev/docker-compose.yml --project-name gobp build {{SERVICES}}



# compile proto files and generate .pb.go files
proto-gen: proto-lint
    # cd src/delivery/grpc && protoc -I=proto --go_out=. --go-grpc_out=. proto/*.proto
    cd src/delivery/grpc/proto && buf generate --template buf.gen.go.yaml
    goimports -l -w src/delivery/grpc/proto/gen/go/

# lint and format .proto files
proto-lint:
    cd src/delivery/grpc/proto && buf lint && buf format -w

# compiles proto fils and generates pb files in python
proto-gen-py:
    cd src/delivery/grpc/proto && buf generate --template buf.gen.py.yaml



# generate mock files for testing using mockery
mockery-gen:
    mockery --dir=src/internal/core/domain/health/ --name=HealthRepoInterface --output=src/internal/repo/testutil/mocks/ --filename=health.repo_mock.go --outpkg=repomock
    mockery --dir=src/internal/core/domain/furniture/ --name=FurnitureRepoInterface --output=src/internal/repo/testutil/mocks/ --filename=furniture.repo_mock.go --outpkg=repomock
    
    mockery --dir=src/internal/core/domain/furniture/ --name=FurnitureUseCaseInterface --output=src/internal/core/usecase/testutil/mocks/ --filename=furniture.uc_mock.go --outpkg=ucmock



# generate schema migration files
migrate-gen:
    just run-docker postgres
    {{GORUN_CMD}} src/cmd/migration/main.go

# apply pending schema migrations to db
migrate-apply:
    atlas migrate apply --url "{{DB_URI}}{{DB_URI_PARAMS}}" --dir "{{MIGRATION_DIR}}"

# show pending migrations
migrate-status:
    atlas migrate status --url "{{DB_URI}}{{DB_URI_PARAMS}}" --dir "{{MIGRATION_DIR}}"

# migrate-set does not actually applies the migration
# mark the migrations upto the given version as applied in revision table
migrate-set version:
    atlas migrate set {{version}} --url "{{DB_URI}}{{DB_URI_PARAMS}}" --dir "{{MIGRATION_DIR}}"

# rollback migrations to specified migration version
migrate-revert version:
    atlas schema apply --url "{{DB_URI}}{{DB_URI_PARAMS}}" --to "{{MIGRATION_DIR}}&version={{version}}" --dev-url "{{DEV_DB_URI}}{{DB_URI_PARAMS}}" --exclude "atlas_schema_revisions"
    just migrate-set {{version}}
    just migrate-show



# output present working directory; useful for debugging recipes
echo-pwd:
    echo "This command is running in '${PWD}'"

# install binaries of go packages needed for development
install-bins:
    ./install-dev-bins.sh
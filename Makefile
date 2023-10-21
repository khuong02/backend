run:
	cd ./cmd/server && go mod tidy && go mod download && \
	go run main.go

migration-up:
	cd cmd/server && go run .\main.go -task migration-up

migration-down:
	cd cmd/server && go run .\main.go -task migration-down

swagger:
	 swag init -g .\cmd\server\main.go -o internal/user/docs \
 		--parseDependency --parseInternal --parseDepth 2

sec:
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec ./...

local:
	docker-compose -f .\docker-compose\docker-compose.yaml --env-file=./docker-compose/env/.env.local up -d

docker-build:
	docker build -t project-test .
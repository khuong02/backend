## DOCUMENT

1. Start DB and enviroment
```bash
make local
or
docker-compose -f .\docker-compose\docker-compose.yaml --env-file=./docker-compose/env/.env.local up -d
```

2. Run migration up or down
* Migrate up
```
make migration-up
or
cd cmd/server && go run .\main.go -task migration-up
```

* Migrate down
```
make migration-down
or
cd cmd/server && go run .\main.go -task migration-down
```

3. Run local
```bash
make run
or
cd ./cmd/server && go mod tidy && go mod download && \
go run main.go
```

4. Gen migration file
```bash
migrate create -ext sql -dir ./internal/user/migrations ${FILE_NAME}
```

5. Swagger
- Install Swag CLI:
```bash
go install github.com/swaggo/swag/cmd/swag
```
- Gen swagger
```bash
make swagger or
swag init -g cmd/server/main.go --parseDependency --parseInternal --parseDepth 2
```

- Swagger Local: http://localhost:8888/docs/index.html

6. Build docker
```bash
make docker-build
or
docker build -t project-test .
```
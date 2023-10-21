run:
	cd ./cmd/server && go mod tidy && go mod download && \
	go run main.go

swagger:
	 swag init -g .\cmd\server\main.go -o internal/user/docs \
 		--parseDependency --parseInternal --parseDepth 2

sec:
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec ./...

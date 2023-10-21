FROM golang:1.20-alpine as modules

RUN apk add --no-cache ca-certificates \
    dpkg \
    gcc \
    git \
    musl-dev \
    openssh

ENV GO111MODULE=on

COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.20-alpine as builder

COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/user

FROM scratch

COPY --from=builder /app/cmd/server/config.yaml ./
COPY --from=builder /app/internal/user/migrations /internal/cms/migrations
COPY --from=builder /bin/app /app

EXPOSE 80

CMD ["/app"]
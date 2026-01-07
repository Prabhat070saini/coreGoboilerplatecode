# -------------------------
# Build stage
# -------------------------
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


# -------------------------
# Runtime stage
# -------------------------
FROM alpine:latest

WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata bash

COPY --from=builder /app/main .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/shared/clients/database/migrations ./migrations
COPY --from=builder /app/config ./config


COPY docker-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh"]

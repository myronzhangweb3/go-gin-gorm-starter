FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o /out/api-server cmd/main.go

FROM alpine:3.18 AS runner

WORKDIR /app

COPY --from=builder /out/* ./

EXPOSE 8000

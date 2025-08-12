FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-app .

EXPOSE 8000

CMD ["./go-app"]

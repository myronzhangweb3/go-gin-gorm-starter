FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o go-gin-gorm-starter main.go

EXPOSE 8081

CMD ["./go-gin-gorm-starter"]

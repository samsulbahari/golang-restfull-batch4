FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/cmd/api cmd/api/main.go

RUN chmod +x /app/cmd/api
ENTRYPOINT ["/app/cmd/api"]





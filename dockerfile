# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOARCH=amd64 go build -o echoserver server.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/echoserver .

RUN chmod +x ./echoserver

EXPOSE 50051

CMD ["./echoserver"]

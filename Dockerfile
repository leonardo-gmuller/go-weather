FROM golang:1.23.9 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY app ./app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goweather ./app/cmd/server

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/goweather .
EXPOSE 8080
CMD ["./goweather"]
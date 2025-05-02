FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod main.go .
RUN go build -o app .
CMD ["./app"]

EXPOSE 1082

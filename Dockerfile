FROM golang:1-alpine

WORKDIR /app
COPY . .

RUN go build -o app .

CMD ["./app"]

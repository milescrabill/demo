FROM golang:1-alpine

WORKDIR /app
COPY . .

RUN go build -o app .
EXPOSE 5000

CMD ["./app"]

FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]
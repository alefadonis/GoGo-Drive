FROM golang:alpine

WORKDIR /gogodrive

COPY . .

RUN go mod download

RUN go build -o main src/*.go

EXPOSE 8081


CMD ["./main"]
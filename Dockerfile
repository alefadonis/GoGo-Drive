FROM golang:alpine

WORKDIR /gogodrive

COPY . .

RUN go mod download

RUN go build src/*.go

EXPOSE 8081


CMD ["./main"]
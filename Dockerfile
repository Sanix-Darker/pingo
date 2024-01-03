FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o pingo

CMD ["./pingo"]

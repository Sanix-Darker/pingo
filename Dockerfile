FROM golang:latest

RUN apt update -y

WORKDIR /src

# FIXME: only for live reload puporse, will be removed on prod.
RUN go install github.com/cosmtrek/air@latest

COPY go.sum go.mod ./
RUN go mod download
RUN go get -d -v ./...
RUN go install -v ./...

COPY . .
RUN cd app && go build -o pingo

WORKDIR /src/app

CMD ["air", "-d", "pingo"]

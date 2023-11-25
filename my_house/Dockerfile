FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && mod verify

COPY .. .
RUN go build -o awesomeProject./main.go

CMD [".house"]

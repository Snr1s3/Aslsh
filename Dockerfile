FROM golang:1.25.5-trixie

WORKDIR /usr/src/app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src/ .

RUN go build -v -o app

CMD ["./app"]
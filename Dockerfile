FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache make

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN go build -o sc-internacional ./cmd

EXPOSE 8080

CMD ["./sc-internacional"]
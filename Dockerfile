FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy
WORKDIR api
RUN go build -o binary .

ENTRYPOINT ["/app/api/binary"]
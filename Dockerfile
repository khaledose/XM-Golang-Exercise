FROM golang:1.19-alpine

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev librdkafka-dev pkgconf

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -tags musl -o ./build/app ./cmd/main.go

EXPOSE 8000

CMD [ "./build/app" ]
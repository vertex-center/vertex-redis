FROM golang:1.19-alpine

ARG PORT=6379

RUN apk --update add redis

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /vertex-redis

EXPOSE $PORT

CMD ["/vertex-redis"]

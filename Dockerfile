FROM golang:1.23.2-alpine

RUN apk add --no-cache \
    bash \
    git \
    gcc \
    musl-dev \
    curl

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# EXPOSE 6666

CMD ["./main"]

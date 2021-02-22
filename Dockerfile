# https://dev.to/plutov/docker-and-go-modules-3kkn
FROM golang:1.15-alpine as builder

RUN set -ex && \
    apk add --no-cache gcc musl-dev

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build

# final stage
FROM alpine:latest
COPY --from=builder /app/grpc-quickstart /app/
EXPOSE 50051

ENTRYPOINT ["/app/grpc-quickstart"]

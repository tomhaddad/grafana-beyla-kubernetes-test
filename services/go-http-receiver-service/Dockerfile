FROM node:20-alpine AS installer
RUN apk add --no-cache git libc6-compat

WORKDIR /app
COPY ./proto/ ./proto
WORKDIR /app/proto
RUN npm install
RUN npm run build

FROM golang:1.20.2-alpine AS builder
RUN apk add --no-cache git protoc protobuf-dev libc6-compat
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR /app
RUN printf "go 1.20 \nuse ./services/go-http-receiver-service\nuse ./proto\n" >> go.work
COPY --from=installer /app/proto/ ./proto

WORKDIR /app/services/go-http-receiver-service/
COPY ./services/go-http-receiver-service ./

RUN go mod download
RUN go build -o ./build

FROM debian:stable-slim AS server
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

WORKDIR /app
COPY --from=builder /app/services/go-http-receiver-service/ .

EXPOSE 8080
CMD [ "./build" ]

FROM golang:1.22

WORKDIR /grpc

COPY go.mod /grpc
COPY go.sum /grpc
COPY /grpc /grpc

RUN go mod download
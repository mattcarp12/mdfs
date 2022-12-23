FROM golang:1.19-buster AS build

WORKDIR /app

# RUN go mod download
RUN go install github.com/grpc-ecosystem/grpc-health-probe@latest
RUN go install github.com/cosmtrek/air@latest

# USER 1000:1000

CMD [ "air", "--build.cmd", "go build -buildvcs=false -o /tmp/app ./cmd/chunkmaster", "--build.bin", "/tmp/app" ]

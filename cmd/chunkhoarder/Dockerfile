FROM golang:1.19-buster AS build

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/grpc-ecosystem/grpc-health-probe@latest

# USER 1000:1000

CMD [ "air", "--build.cmd", "go build -buildvcs=false -o /tmp/app ./cmd/chunkhoarder", "--build.bin", "/tmp/app" ]

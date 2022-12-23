pb:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./grpc/mdfs.proto


# master:
#     go run ./cmd/chunkmaster

# hoarder:
#     go run ./cmd/chunkhoarder
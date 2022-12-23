package chunkmaster

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/mattcarp12/mdfs/grpc"
	"github.com/mattcarp12/mdfs/lib/kv"
	"google.golang.org/grpc"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedChunkMasterServiceServer

	kv *kv.KV
}

func NewChunkmasterServer(ctx context.Context, kv *kv.KV) *Server {
	return &Server{
		kv: kv,
	}
}

const chunkMasterPort = 8845

func (s *Server) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", chunkMasterPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChunkMasterServiceServer(grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) List(ctx context.Context, r *pb.ListRequest) (*pb.ListResponse, error) {
	// Get list of files from redis with matching prefix

	keyList, err := s.kv.ListPrefix(ctx, r.GetPrefix())
	if err != nil {
		return nil, err
	}

	var files []*pb.FileMetadata

	for _, key := range keyList {
		files = append(files, &pb.FileMetadata{Name: key})
	}

	return &pb.ListResponse{Files: files}, nil
}

func (s *Server) GetFileChunks(ctx context.Context, r *pb.File) (*pb.Chunks, error) {

	// Get list of chunks in the file

	// For each chunk, get list of Chunkhoarders that
	// are hoarding that chunk

	// For each replicating Chunkhoarder, load balance
	// the request

	return nil, nil
}

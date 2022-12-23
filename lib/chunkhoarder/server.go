package chunkhoarder

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/mattcarp12/mdfs/grpc"
	"github.com/mattcarp12/mdfs/lib/config"
	"github.com/mattcarp12/mdfs/lib/util"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChunkHoarderServiceServer

	chunkdir string
}

const chunkHoarderPort = 8847

func (s *Server) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", chunkHoarderPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChunkHoarderServiceServer(grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) GetChunk(ctx context.Context, r *pb.GetChunkRequest) (*pb.GetChunkResponse, error) {
	return nil, nil
}

func (s *Server) ListChunks(ctx context.Context, e *pb.Empty) (*pb.Chunks, error) {
	files, err := os.ReadDir(s.chunkdir)
	util.Check(err)

	chunks := filesToChunks(files)

	return chunks, nil
}

func filesToChunks(files []os.DirEntry) *pb.Chunks {
	chunks := pb.Chunks{}
	for _, file := range files {
		chunks.Chunks = append(chunks.Chunks, &pb.Chunk{
			ChunkID: file.Name(),
		})
	}

	return &chunks
}

func (s *Server) PutChunk(ctx context.Context, r *pb.PutChunkRequest) (*pb.PutChunkResponse, error) {
	return nil, nil
}

func NewChunkhoarderServer(ctx context.Context, config config.Config) *Server {
	// check if chunkdir exists
	if _, err := os.Stat(config.ChunkDir); os.IsNotExist(err) {
		// try to create it
		err := os.MkdirAll(config.ChunkDir, 0666)
		util.Check(err)
	}
	return &Server{
		chunkdir: config.ChunkDir,
	}
}

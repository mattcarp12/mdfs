package chunkmaster

import (
	"context"
	"log"

	pb "github.com/mattcarp12/mdfs/grpc"
	"github.com/mattcarp12/mdfs/lib/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChunkmasterClient struct {
	conn   *grpc.ClientConn
	client pb.ChunkMasterServiceClient
	clock  uint64
}

func NewChunkmasterClient(addr string) *ChunkmasterClient {
	// Set up a connection to the server.
	// TODO : Make secure connection.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewChunkMasterServiceClient(conn)

	return &ChunkmasterClient{
		conn:   conn,
		client: c,
	}
}

func (c *ChunkmasterClient) List(ctx context.Context, prefix string) ([]string, error) {
	resp, err := c.client.List(ctx, &pb.ListRequest{Prefix: prefix})
	if err != nil {
		return nil, err
	}

	var fileList []string
	for _, f := range resp.Files {
		fileList = append(fileList, f.GetName())
	}

	return fileList, nil
}

func (c *ChunkmasterClient) Heartbeat(ctx context.Context) error {
	// My name is what?
	name := util.Hostname()
	c.clock++

	req := &pb.HeartbeatRequest{
		Name:  name,
		Clock: c.clock,
		FsMeta: &pb.FsMeta{
			FreeBytes: 0,
			NumChunks: 0,
		},
	}

	resp, err := c.client.Heartbeat(ctx, req)
	// What to do if there is an error here?
	log.Printf("Heartbeat Response: %s\n", resp)

	return err
}

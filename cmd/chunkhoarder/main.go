package main

import (
	"context"
	"log"

	"github.com/mattcarp12/mdfs/lib/chunkhoarder"
	"github.com/mattcarp12/mdfs/lib/chunkmaster"
	"github.com/mattcarp12/mdfs/lib/config"
)

func main() {
	ctx := context.Background()
	config := config.NewConfig()

	log.Printf("Context: %+v, Config: %+v\n", ctx, config)

	chunkMasterClient := chunkmaster.NewChunkmasterClient(config.ChunkmasterAddr)
	heartbeater := chunkhoarder.NewHeartbeater(chunkMasterClient)
	heartbeater.StartBeating(ctx)

	// Start the chunkhoarder server
	server := chunkhoarder.NewChunkhoarderServer(ctx, config)
	server.Start()
}

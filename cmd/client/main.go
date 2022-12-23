package main

import (
	"context"
	"log"
	"time"

	"github.com/mattcarp12/mdfs/lib/chunkmaster"
	"github.com/mattcarp12/mdfs/lib/config"
	"github.com/mattcarp12/mdfs/lib/util"
)

func main() {
	config := config.NewConfig()
	StartClient(config)
}

func StartClient(config config.Config) {

	chunkMasterClient := chunkmaster.NewChunkmasterClient(config.ChunkmasterAddr)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get list of files
	ls, err := chunkMasterClient.List(ctx, "/")
	util.Check(err)

	log.Printf("List of files: %+v\n", ls)
}

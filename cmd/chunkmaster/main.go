package main

import (
	"context"

	"github.com/mattcarp12/mdfs/lib/chunkmaster"
	"github.com/mattcarp12/mdfs/lib/config"
	"github.com/mattcarp12/mdfs/lib/kv"
)

func main() {
	ctx := context.Background()

	config := config.NewConfig()

	kv := kv.NewKV(ctx, config.RedisAddr)

	server := chunkmaster.NewChunkmasterServer(ctx, kv)

	server.Start()
}

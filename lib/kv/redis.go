package kv

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/mattcarp12/mdfs/lib/util"
)

type KV struct {
	redis *redis.Client
}

func NewKV(ctx context.Context, addr string) *KV {
	client := redis.NewClient(&redis.Options{Addr: addr})
	_, err := client.Ping(ctx).Result()
	util.Check(err)

	return &KV{
		redis: client,
	}
}

func (kv *KV) ListPrefix(ctx context.Context, prefix string) ([]string, error) {
	log.Printf("KV: %+v\n", kv)

	// time.Sleep(3 * time.Second)

	keys, err := kv.redis.Keys(ctx, prefix).Result()
	if err != nil {
		log.Printf("ERROR: %+v", err)
	}

	return keys, err
}

func (kv *KV) Redis() *redis.Client {
	return kv.redis
}

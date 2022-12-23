package chunkhoarder

import (
	"context"
	"time"

	"github.com/mattcarp12/mdfs/lib/chunkmaster"
	"github.com/mattcarp12/mdfs/lib/util"
)

type Heartbeater struct {
	client *chunkmaster.ChunkmasterClient
}

func NewHeartbeater(client *chunkmaster.ChunkmasterClient) *Heartbeater {
	return &Heartbeater{client}
}

func (h *Heartbeater) Do(ctx context.Context) {
	err := h.client.Heartbeat(ctx)
	util.Check(err)
}

func (h *Heartbeater) StartBeating(ctx context.Context) {
	go func() {
		for {
			h.Do(ctx)
			time.Sleep(10 * time.Second)
		}
	}()
}

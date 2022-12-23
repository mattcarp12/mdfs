package chunkmaster

import (
	"context"
	"log"
	"strconv"

	pb "github.com/mattcarp12/mdfs/grpc"
	"github.com/mattcarp12/mdfs/lib/util"
)

func (s *Server) Heartbeat(ctx context.Context, r *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	logHeartbeat(r)

	// When the server gets a hearbeat what does it do?
	// It updates the entry in it's database (redis).
	// What is the key? chunkhoarder:<name> = hashmap data structure
	s.heartbeatUpdateKV(ctx, r)

	// After the KV is updated, what needs to happen? Need to make the response.
	// What is the response? We'll need to research more about how other DFSs
	// do their heartbeat. What do we want out of the heartbeat? Synchronization?
	// The master needs to know how much space the chunkhoards have so it can
	// properly load balance. So it needs the load balance data when responding to
	// PutFile requests. It may also want other data to load balance GetFile requests.
	valStr, err := s.kv.Redis().Get(ctx, chunkhoarderClock(r.GetName())).Result()
	util.Check(err)

	val, err := strconv.ParseUint(valStr, 10, 64)
	util.Check(err)

	return &pb.HeartbeatResponse{Clock: val}, nil
}

func (s *Server) heartbeatUpdateKV(ctx context.Context, r *pb.HeartbeatRequest) error {
	s.kv.Redis().Set(ctx, chunkhoarderClock(r.GetName()), r.GetClock(), 0)
	s.kv.Redis().Set(ctx, chunkhoarderKey(r.GetName())+":freeBytes", r.GetFsMeta().GetFreeBytes(), 0)
	// s.kv.Redis().Set(ctx, chunkhoarderKey(r.GetName())+":freeBytes", r.GetFsMeta().GetFreeBytes(), 0)

	return nil
}

func chunkhoarderKey(name string) string {
	return "chunkhoarder:" + name
}

func chunkhoarderClock(name string) string {
	return chunkhoarderKey(name) + ":clock"
}

func logHeartbeat(r *pb.HeartbeatRequest) {
	log.Printf(
		"Received heartbeat from %s. Clock: %d. FreeSpace: %d. NumChunks: %d",
		r.GetName(),
		r.GetClock(),
		r.GetFsMeta().GetFreeBytes(),
		r.GetFsMeta().GetNumChunks(),
	)
}

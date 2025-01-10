package grpc

import (
	"context"
	"time"

	goginboilerplate "github.com/dezh-tech/geb"
	pb "github.com/dezh-tech/geb/delivery/grpc/gen"
)

type healthServer struct {
	*Server
}

func newHealthServer(server *Server) *healthServer {
	return &healthServer{
		Server: server,
	}
}

func (s healthServer) Status(ctx context.Context, _ *pb.StatusRequest) (*pb.StatusResponse, error) {
	services := make([]*pb.Service, 0)

	redisStatus := pb.Status_CONNECTED
	redisMessage := ""

	if err := s.Redis.Client.Ping(ctx).Err(); err != nil {
		redisStatus = pb.Status_DISCONNECTED
		redisMessage = err.Error()
	}

	redis := pb.Service{
		Name:    "redis",
		Status:  redisStatus,
		Message: redisMessage,
	}

	services = append(services, &redis)

	mongoStatus := pb.Status_CONNECTED
	mongoMessage := ""

	if err := s.DB.Client.Ping(ctx, nil); err != nil {
		mongoStatus = pb.Status_DISCONNECTED
		mongoMessage = err.Error()
	}

	mongo := pb.Service{
		Name:    "mongo",
		Status:  mongoStatus,
		Message: mongoMessage,
	}

	services = append(services, &mongo)

	return &pb.StatusResponse{
		Uptime:   int64(time.Since(s.StartTime).Seconds()),
		Version:  goginboilerplate.StringVersion(),
		Services: services,
	}, nil
}

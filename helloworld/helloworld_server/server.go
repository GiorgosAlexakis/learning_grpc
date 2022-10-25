package main

import (
	"context"

	pb "github.com/Looper2074/learning_grpc/helloworld/helloworld"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	// simulating token validation
	if r.Token == 0 {
		return &pb.HelloResponse{}, status.Errorf(codes.PermissionDenied, "invalid token")
	}
	return &pb.HelloResponse{Message: "Hello " + r.GetName()}, nil
}
func main() {

}

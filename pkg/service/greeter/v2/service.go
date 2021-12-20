package v2

import (
	"context"

	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	db "github.com/danlmarmot/grpc-v-greeter/pkg/database"
	"github.com/danlmarmot/grpc-v-greeter/pkg/serializer"
	pb "github.com/danlmarmot/grpc-v-greeter/protogen/v2/greeter"
)

type server struct {
	ctx        context.Context
	serializer serializer.DatabaseSerializer
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	criteria := db.RegisterDetailsCriteria{
		First: req.FirstName,
		Last:  req.LastName,
	}
	log.Printf("processing SayHello request: %v", req)

	resp, err := s.serializer.RegisterDetails(ctx, criteria)
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to register user")
	}
	return resp.ToV2()
}

func (s *server) SayGoodbye(ctx context.Context, req *pb.SayGoodbyeRequest) (*pb.SayGoodbyeResponse, error) {
	criteria := db.DeleteUserCriteria{
		First: req.FirstName,
		Last:  req.LastName,
	}
	log.Printf("processing SayGoodbye request: %v", req)

	resp, err := s.serializer.DeleteUser(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return resp.ToV2()
}

func NewService() (*server, error) {
	sz, err := serializer.NewDBSerializer()
	if err != nil {
		return nil, err
	}
	s := &server{
		ctx:        context.Background(),
		serializer: sz,
	}
	return s, nil
}

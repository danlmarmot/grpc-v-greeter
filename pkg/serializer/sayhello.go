package serializer

import (
	"context"
	"fmt"
	db "github.com/danlmarmot/grpc-v-greeter/pkg/database"
	v1 "github.com/danlmarmot/grpc-v-greeter/protogen/v1/greeter"
	v2 "github.com/danlmarmot/grpc-v-greeter/protogen/v2/greeter"
)

type registerDetailsWrapper struct {
	*db.RegisterDetailsResp
}

func (w *registerDetailsWrapper) ToV1() (*v1.SayHelloResponse, error) {
	if w.Exists {
		return &v1.SayHelloResponse{
			Message: fmt.Sprintf("Welcome back %s", w.User.FirstName),
		}, nil
	}
	return &v1.SayHelloResponse{
		Message: fmt.Sprintf("Nice to meet you %s", w.User.FirstName),
	}, nil
}

func (w *registerDetailsWrapper) ToV2() (*v2.SayHelloResponse, error) {
	if w.Exists {
		return &v2.SayHelloResponse{
			Message: fmt.Sprintf("Welcome back %s %s", w.User.FirstName, w.User.LastName),
		}, nil
	}
	return &v2.SayHelloResponse{
		Message: fmt.Sprintf("Nice to meet you %s %s", w.User.FirstName, w.User.LastName),
	}, nil
}

func (dbs *dbserializer) RegisterDetails(ctx context.Context, criteria db.RegisterDetailsCriteria) (SayHelloSerializer, error) {
	resp, err := dbs.storer.RegisterDetails(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return &registerDetailsWrapper{resp}, nil
}

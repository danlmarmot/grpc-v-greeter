package serializer

import (
	"context"
	"fmt"
	db "github.com/danlmarmot/grpc-v-greeter/pkg/database"
	v1 "github.com/danlmarmot/grpc-v-greeter/protogen/v1/greeter"
	v2 "github.com/danlmarmot/grpc-v-greeter/protogen/v2/greeter"
)

type deleteUserWrapper struct {
	*db.DeleteUserResp
}

func (w *deleteUserWrapper) ToV1() (*v1.SayGoodbyeResponse, error) {
	return &v1.SayGoodbyeResponse{
		Message: fmt.Sprintf("Sad to see you go %s, you visited %d times", w.User.FirstName, w.Count),
	}, nil
}

func (w *deleteUserWrapper) ToV2() (*v2.SayGoodbyeResponse, error) {
	return &v2.SayGoodbyeResponse{
		Message: fmt.Sprintf("Sad to see you go %s %s, you visited %d times",
			w.User.FirstName, w.User.LastName, w.Count),
	}, nil
}

func (dbs *dbserializer) DeleteUser(ctx context.Context, criteria db.DeleteUserCriteria) (SayGoodbyeSerializer, error) {
	resp, err := dbs.storer.DeleteUser(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return &deleteUserWrapper{resp}, nil
}

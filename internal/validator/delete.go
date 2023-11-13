package validator

import (
	"context"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *implementation) PlayerDelete(ctx context.Context, in *pb.PlayerDeleteRequest) (*emptypb.Empty, error) {
	method := "delete request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	if in.Id < 1 {
		msg := "id should be greater than 0"
		i.log.Errorf("%s: %s", method, msg)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	response, err := i.client.PlayerDelete(ctx, in)
	counter.PushOutRequestsCounter()

	if err != nil {
		i.log.Errorf("%s: client error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return nil, err
	}

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	if err = i.redis.Delete(i.redisKey(in.Id)); err != nil {
		i.log.Errorf("%s: redis delete error: %s", method, err)
	}

	return response, err
}

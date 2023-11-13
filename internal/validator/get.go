package validator

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (i *implementation) PlayerGet(ctx context.Context, in *pb.PlayerGetRequest) (*pb.PlayerGetResponse, error) {
	method := "get request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	if in.Id < 1 {
		msg := "id should be greater than 0"
		i.log.Errorf("%s: %s", method, msg)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	key := fmt.Sprintf("player:%d", in.Id)
	var value *pb.PlayerGetResponse

	err := i.redis.Get(key, &value)
	if err != nil {
		i.log.Errorf("%s: redis error: %s", method, err)
		counter.PushMissRedisCounter()
	} else {
		i.log.Infof("%s: Done! (redis response)", method)
		counter.PushSuccessRequestsCounter()
		counter.PushHitRedisCounter()
		return value, nil
	}

	response, err := i.client.PlayerGet(ctx, in)
	counter.PushOutRequestsCounter()

	if err != nil {
		i.log.Errorf("%s: client error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return nil, err
	}

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	if err = i.redis.Set(key, response, time.Minute); err != nil {
		i.log.Errorf("%s: redis set error: %s", method, err)
	}

	return response, err
}

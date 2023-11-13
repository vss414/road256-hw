package validator

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"time"
)

func (i *implementation) PlayerList(ctx context.Context, in *pb.PlayerListRequest) (*pb.PlayerListResponse, error) {
	method := "list request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	key := fmt.Sprintf(
		"list:page:%d:limit:%d:order:%d:direction:%d",
		in.GetPage(),
		in.GetLimit(),
		in.GetOrder(),
		in.GetDirection(),
	)
	var value *pb.PlayerListResponse
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

	response, err := i.client.PlayerList(ctx, in)
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

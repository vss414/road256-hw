package validator

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	playerModelPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (i *implementation) PlayerCreate(ctx context.Context, in *pb.PlayerCreateRequest) (*pb.PlayerCreateResponse, error) {
	method := "create request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	_, err := playerModelPkg.New(
		in.GetName(),
		in.GetClub(),
		uint(in.GetGames()),
		uint(in.GetGoals()),
		uint(in.GetAssists()),
	)

	if err != nil {
		if errors.Is(err, playerModelPkg.ErrValidation) {
			i.log.Errorf("%s: validation error: %s", method, err)
			counter.PushFailedRequestsCounter()
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		i.log.Errorf("%s: internal error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.Internal, err.Error())
	}

	response, err := i.client.PlayerCreate(ctx, in)
	counter.PushOutRequestsCounter()

	if err != nil {
		i.log.Errorf("%s: client error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return nil, err
	}

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	if err = i.redis.Set(i.redisKey(response.Id), response, time.Minute); err != nil {
		i.log.Errorf("%s: redis set error: %s", method, err)
	}

	return response, err
}

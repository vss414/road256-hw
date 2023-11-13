package api

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) PlayerGet(ctx context.Context, in *pb.PlayerGetRequest) (*pb.PlayerGetResponse, error) {
	method := "get request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	player, err := i.player.Get(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, repository.ErrPlayerNotExists) {
			i.log.Errorf("%s: validation error: %s", method, err)
			counter.PushFailedRequestsCounter()
			return nil, status.Error(codes.NotFound, err.Error())
		}

		i.log.Errorf("%s: internal error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.Internal, err.Error())
	}

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	return &pb.PlayerGetResponse{
		Id:      player.Id,
		Name:    player.Name,
		Club:    player.Club,
		Games:   uint64(player.Games),
		Goals:   uint64(player.Goals),
		Assists: uint64(player.Assists),
	}, nil
}

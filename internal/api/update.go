package api

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	playerModelPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *implementation) PlayerUpdate(ctx context.Context, in *pb.PlayerUpdateRequest) (*emptypb.Empty, error) {
	method := "update request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	p := playerModelPkg.Player{
		Id:      in.GetId(),
		Name:    in.GetName(),
		Club:    in.GetClub(),
		Games:   uint(in.GetGames()),
		Goals:   uint(in.GetGoals()),
		Assists: uint(in.GetAssists()),
	}

	if err := i.player.Update(ctx, &p); err != nil {
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

	return &emptypb.Empty{}, nil
}

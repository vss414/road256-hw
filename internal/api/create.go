package api

import (
	"context"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	playerModelPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) PlayerCreate(ctx context.Context, in *pb.PlayerCreateRequest) (*pb.PlayerCreateResponse, error) {
	method := "create request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	p := playerModelPkg.Player{
		Name:    in.GetName(),
		Club:    in.GetClub(),
		Games:   uint(in.GetGames()),
		Goals:   uint(in.GetGoals()),
		Assists: uint(in.GetAssists()),
	}

	player, err := i.player.Create(ctx, &p)
	if err != nil {
		i.log.Errorf("%s: internal error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.Internal, err.Error())
	}

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	return &pb.PlayerCreateResponse{
		Id:      player.Id,
		Name:    player.Name,
		Club:    player.Club,
		Games:   uint64(player.Games),
		Goals:   uint64(player.Goals),
		Assists: uint64(player.Assists),
	}, nil
}

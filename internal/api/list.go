package api

import (
	"context"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) PlayerList(ctx context.Context, in *pb.PlayerListRequest) (*pb.PlayerListResponse, error) {
	method := "list request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	players, err := i.player.List(ctx, in.GetLimit(), in.GetPage(), in.GetOrder().String(), in.GetDirection().String())
	if err != nil {
		i.log.Errorf("%s: internal error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return &pb.PlayerListResponse{}, status.Error(codes.Internal, err.Error())
	}

	result := make([]*pb.PlayerListResponse_Player, 0, len(players))
	for _, player := range players {
		result = append(result, &pb.PlayerListResponse_Player{
			Id:      player.Id,
			Name:    player.Name,
			Club:    player.Club,
			Games:   uint64(player.Games),
			Goals:   uint64(player.Goals),
			Assists: uint64(player.Assists),
		})
	}

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	return &pb.PlayerListResponse{
		Players: result,
	}, nil
}

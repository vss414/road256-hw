package api

import (
	"context"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

func (i *implementation) PlayerStreamList(in *pb.PlayerStreamListRequest, srv pb.Admin_PlayerStreamListServer) error {
	method := "stream list request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	ctx := context.Background()

	players, err := i.player.List(ctx, in.GetLimit(), in.GetPage(), in.GetOrder().String(), in.GetDirection().String())
	if err != nil {
		i.log.Errorf("%s: internal error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return status.Error(codes.Internal, err.Error())
	}

	var wg sync.WaitGroup
	for _, player := range players {
		wg.Add(1)
		go func(p *model.Player) {
			defer wg.Done()

			if err := srv.Send(&pb.PlayerStreamListResponse{
				Id:      p.Id,
				Name:    p.Name,
				Club:    p.Club,
				Games:   uint64(p.Games),
				Goals:   uint64(p.Goals),
				Assists: uint64(p.Assists),
			}); err != nil {
				i.log.Errorf("%s: send error: %s", method, err)
				log.Printf("send error %v", err)
			}
		}(player)
	}

	wg.Wait()

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	return nil
}

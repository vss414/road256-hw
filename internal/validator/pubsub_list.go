package validator

import (
	"context"
	"encoding/json"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	redis_consumer "gitlab.ozon.dev/vss414/hw-1/internal/redis-consumer"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const listMethod = "list async request"

func (i *implementation) PlayerPubsubList(ctx context.Context, in *pb.PlayerListRequest) (*pb.PlayerListResponse, error) {
	i.log.Infof("%s: %v", listMethod, in)
	counter.PushInRequestsCounter()

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	ch := make(chan *pb.PlayerListResponse)
	errCh := make(chan error)

	go i.pubsubList(in, ch, errCh)

	select {
	case p := <-ch:
		return p, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (i *implementation) pubsubList(in *pb.PlayerListRequest, ch chan *pb.PlayerListResponse, errCh chan error) {
	i.log.Infof("%s: marshal message", listMethod)
	redisMessage, err := json.Marshal(in)
	if err != nil {
		i.log.Errorf("%s: marshal error: %s", listMethod, err)
		counter.PushFailedRequestsCounter()
		errCh <- status.Error(codes.Internal, err.Error())
	}

	go func() {
		i.log.Infof("%s: publish message", listMethod)
		if err := i.redis.Publish(redis_consumer.PlayerList, redisMessage); err != nil {
			i.log.Errorf("%s: redis publish error: %s", listMethod, err)
			errCh <- status.Error(codes.Internal, err.Error())
		}
	}()

	i.log.Infof("%s: subscribe", listMethod)
	pubsub := i.redis.Subscribe(redis_consumer.PlayerListResponse, redis_consumer.PlayerListError)
	defer pubsub.Close()

	i.log.Infof("%s: receive message", listMethod)
	msg, err := pubsub.ReceiveMessage()
	if err != nil {
		i.log.Errorf("%s: receive message error: %s", listMethod, err)
		errCh <- status.Error(codes.Internal, err.Error())
	}

	i.log.Infof("%s: unmarshal message", listMethod)
	switch msg.Channel {
	case redis_consumer.PlayerListResponse:
		var players []*model.Player
		if err := json.Unmarshal([]byte(msg.Payload), &players); err != nil {
			i.log.Errorf("%s: unmarshal message error: %s", listMethod, err)
			errCh <- status.Error(codes.Internal, err.Error())
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

		i.log.Infof("%s: Done!", listMethod)
		counter.PushSuccessRequestsCounter()

		ch <- &pb.PlayerListResponse{
			Players: result,
		}
	case redis_consumer.PlayerListError:
		i.log.Errorf("%s: internal error: %s", listMethod, msg.Payload)
		errCh <- status.Error(codes.Internal, msg.Payload)
	}
}

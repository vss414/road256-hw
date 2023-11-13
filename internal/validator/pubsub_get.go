package validator

import (
	"context"
	"encoding/json"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	redis_consumer "gitlab.ozon.dev/vss414/hw-1/internal/redis-consumer"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

const getMethod = "get async request"

func (i *implementation) PlayerPubsubGet(ctx context.Context, in *pb.PlayerGetRequest) (*pb.PlayerGetResponse, error) {
	i.log.Infof("%s: %v", getMethod, in)
	counter.PushInRequestsCounter()

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	if in.Id < 1 {
		msg := "id should be greater than 0"
		i.log.Errorf("%s: %s", getMethod, msg)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	ch := make(chan *pb.PlayerGetResponse)
	errCh := make(chan error)

	go i.pubsubGet(in, ch, errCh)

	select {
	case p := <-ch:
		return p, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (i *implementation) pubsubGet(in *pb.PlayerGetRequest, ch chan *pb.PlayerGetResponse, errCh chan error) {
	i.log.Infof("%s: marshal message", getMethod)
	redisMessage, err := json.Marshal(in)
	if err != nil {
		i.log.Errorf("%s: marshal error: %s", getMethod, err)
		counter.PushFailedRequestsCounter()
		errCh <- status.Error(codes.Internal, err.Error())
	}

	go func() {
		i.log.Infof("%s: publish message", getMethod)
		if err := i.redis.Publish(redis_consumer.PlayerGet, redisMessage); err != nil {
			i.log.Errorf("%s: redis publish error: %s", getMethod, err)
			errCh <- status.Error(codes.Internal, err.Error())
		}
	}()

	i.log.Infof("%s: subscribe", getMethod)
	pubsub := i.redis.Subscribe(redis_consumer.PlayerGetResponse, redis_consumer.PlayerGetError)
	defer pubsub.Close()

	i.log.Infof("%s: receive message", getMethod)
	msg, err := pubsub.ReceiveMessage()
	if err != nil {
		i.log.Errorf("%s: receive message error: %s", getMethod, err)
		errCh <- status.Error(codes.Internal, err.Error())
	}

	i.log.Infof("%s: unmarshal message", getMethod)
	switch msg.Channel {
	case redis_consumer.PlayerGetResponse:
		var response *pb.PlayerGetResponse
		if err := json.Unmarshal([]byte(msg.Payload), &response); err != nil {
			i.log.Errorf("%s: unmarshal message error: %s", getMethod, err)
			errCh <- status.Error(codes.Internal, err.Error())
		}

		i.log.Infof("%s: Done!", getMethod)
		counter.PushSuccessRequestsCounter()
		ch <- response
	case redis_consumer.PlayerGetError:
		if strings.Contains(msg.Payload, repository.ErrPlayerNotExists.Error()) {
			i.log.Errorf("%s: player not found", getMethod)
			errCh <- status.Error(codes.NotFound, msg.Payload)
		} else {
			i.log.Errorf("%s: internal error: %s", getMethod, msg.Payload)
			errCh <- status.Error(codes.Internal, msg.Payload)
		}
	}
}

package validator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/consumer"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	playerModelPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *implementation) PlayerAsyncUpdate(ctx context.Context, in *pb.PlayerUpdateRequest) (*emptypb.Empty, error) {
	method := "async update request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	if in.Id < 1 {
		i.log.Errorf("%s: id should be greater than 0", method)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.InvalidArgument, "id should be greater than 0")
	}

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

	msg, err := json.Marshal(in)
	if err != nil {
		i.log.Errorf("%s: marshal error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.Internal, err.Error())
	}

	i.producer.Input() <- &sarama.ProducerMessage{
		Topic: consumer.PlayerEventsTopic,
		Key:   sarama.ByteEncoder(fmt.Sprintf("update-%d", in.Id)),
		Value: sarama.ByteEncoder(msg),
	}

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	return &emptypb.Empty{}, nil
}

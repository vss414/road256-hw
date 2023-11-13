package validator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/vss414/hw-1/internal/consumer"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *implementation) PlayerAsyncDelete(ctx context.Context, in *pb.PlayerDeleteRequest) (*emptypb.Empty, error) {
	method := "async delete request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	if in.Id < 1 {
		msg := "id should be greater than 0"
		i.log.Errorf("%s: %s", method, msg)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	msg, err := json.Marshal(in)
	if err != nil {
		i.log.Errorf("%s: marshal error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return nil, status.Error(codes.Internal, err.Error())
	}

	i.producer.Input() <- &sarama.ProducerMessage{
		Topic: consumer.PlayerEventsTopic,
		Key:   sarama.ByteEncoder(fmt.Sprintf("delete-%d", in.Id)),
		Value: sarama.ByteEncoder(msg),
	}

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	return &emptypb.Empty{}, nil
}

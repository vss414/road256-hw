package api

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc"
	"testing"
)

func makeStreamMock() *StreamMock {
	return &StreamMock{
		ctx:            context.Background(),
		sentFromServer: make(chan *pb.PlayerStreamListResponse, 1),
	}
}

type StreamMock struct {
	grpc.ServerStream
	ctx            context.Context
	sentFromServer chan *pb.PlayerStreamListResponse
}

func (m *StreamMock) Send(resp *pb.PlayerStreamListResponse) error {
	m.sentFromServer <- resp
	return nil
}

func TestStream(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		stream := makeStreamMock()
		player := &model.Player{
			Id:      10,
			Name:    "Messi",
			Club:    "PSG",
			Games:   546,
			Goals:   480,
			Assists: 197,
		}

		f.repository.EXPECT().
			List(f.ctx, uint64(3), uint64(1), "ORDER_ID", "DIRECTION_ASC").
			Times(1).
			Return([]*model.Player{player}, nil)

		err := f.service.PlayerStreamList(&pb.PlayerStreamListRequest{
			Limit:     3,
			Page:      1,
			Order:     1,
			Direction: 1,
		}, stream)

		require.NoError(t, err)

		done := make(chan bool)
		go func() {
			for {
				select {
				case p := <-stream.sentFromServer:
					assert.IsType(t, &pb.PlayerStreamListResponse{}, p)
					assert.Equal(t, &player.Id, &p.Id)
					assert.Equal(t, player.Name, p.Name)
					assert.Equal(t, player.Club, p.Club)
					assert.Equal(t, uint64(player.Games), p.Games)
					assert.Equal(t, uint64(player.Goals), p.Goals)
					assert.Equal(t, uint64(player.Assists), p.Assists)
				default:
					done <- true
					return
				}
			}
		}()
		<-done
	})

	t.Run("negative", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		stream := makeStreamMock()
		expectedError := errors.New("some error")

		f.repository.EXPECT().
			List(f.ctx, uint64(3), uint64(1), "ORDER_ID", "DIRECTION_ASC").
			Times(1).
			Return(nil, expectedError)

		err := f.service.PlayerStreamList(&pb.PlayerStreamListRequest{
			Limit:     3,
			Page:      1,
			Order:     1,
			Direction: 1,
		}, stream)

		assert.Equal(t, "rpc error: code = Internal desc = some error", err.Error())
	})
}

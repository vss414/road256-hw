package api

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

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

		response, err := f.service.PlayerList(f.ctx, &pb.PlayerListRequest{
			Limit:     3,
			Page:      1,
			Order:     pb.Order_ORDER_ID,
			Direction: pb.Direction_DIRECTION_ASC,
		})

		require.NoError(t, err)
		assert.Len(t, response.Players, 1)
		assert.Equal(t, player.Id, response.Players[0].Id)
		assert.Equal(t, player.Name, response.Players[0].Name)
		assert.Equal(t, player.Club, response.Players[0].Club)
		assert.Equal(t, uint64(player.Games), response.Players[0].Games)
		assert.Equal(t, uint64(player.Goals), response.Players[0].Goals)
		assert.Equal(t, uint64(player.Assists), response.Players[0].Assists)
	})

	t.Run("negative", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		expectedError := errors.New("some error")

		f.repository.EXPECT().
			List(f.ctx, uint64(3), uint64(1), "ORDER_ID", "DIRECTION_ASC").
			Times(1).
			Return(nil, expectedError)

		response, err := f.service.PlayerList(f.ctx, &pb.PlayerListRequest{
			Limit:     3,
			Page:      1,
			Order:     pb.Order_ORDER_ID,
			Direction: pb.Direction_DIRECTION_ASC,
		})

		assert.Equal(t, &pb.PlayerListResponse{}, response)
		assert.Equal(t, "rpc error: code = Internal desc = some error", err.Error())
	})
}

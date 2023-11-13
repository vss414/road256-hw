package api

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		request := &model.Player{
			Name:    "Messi",
			Club:    "PSG",
			Games:   546,
			Goals:   480,
			Assists: 197,
		}
		player := &model.Player{
			Id:      10,
			Name:    request.Name,
			Club:    request.Club,
			Games:   request.Games,
			Goals:   request.Goals,
			Assists: request.Assists,
		}

		f.repository.EXPECT().Create(f.ctx, request).Times(1).Return(player, nil)

		response, err := f.service.PlayerCreate(f.ctx, &pb.PlayerCreateRequest{
			Name:    player.Name,
			Club:    player.Club,
			Games:   uint64(player.Games),
			Goals:   uint64(player.Goals),
			Assists: uint64(player.Assists),
		})

		require.NoError(t, err)
		assert.IsType(t, &pb.PlayerCreateResponse{}, response)
		assert.Equal(t, player.Id, response.Id)
		assert.Equal(t, player.Name, response.Name)
		assert.Equal(t, player.Club, response.Club)
		assert.Equal(t, uint64(player.Games), response.Games)
		assert.Equal(t, uint64(player.Goals), response.Goals)
		assert.Equal(t, uint64(player.Assists), response.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		expectedError := errors.New("some error")
		player := &model.Player{
			Name:    "Messi",
			Club:    "PSG",
			Games:   546,
			Goals:   480,
			Assists: 197,
		}

		f.repository.EXPECT().Create(f.ctx, player).Times(1).Return(nil, expectedError)

		response, err := f.service.PlayerCreate(f.ctx, &pb.PlayerCreateRequest{
			Name:    player.Name,
			Club:    player.Club,
			Games:   uint64(player.Games),
			Goals:   uint64(player.Goals),
			Assists: uint64(player.Assists),
		})

		assert.Nil(t, response)
		assert.Equal(t, "rpc error: code = Internal desc = some error", err.Error())
	})
}

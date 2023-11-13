package api

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGet(t *testing.T) {
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

		f.repository.EXPECT().Get(f.ctx, player.Id).Times(1).Return(player, nil)

		response, err := f.service.PlayerGet(f.ctx, &pb.PlayerGetRequest{Id: player.Id})

		require.NoError(t, err)
		assert.IsType(t, &pb.PlayerGetResponse{}, response)
		assert.Equal(t, player.Id, response.Id)
		assert.Equal(t, player.Name, response.Name)
		assert.Equal(t, player.Club, response.Club)
		assert.Equal(t, uint64(player.Games), response.Games)
		assert.Equal(t, uint64(player.Goals), response.Goals)
		assert.Equal(t, uint64(player.Assists), response.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("player not found", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			var id uint64 = 10
			f.repository.EXPECT().Get(f.ctx, id).Times(1).Return(nil, repository.ErrPlayerNotExists)

			response, err := f.service.PlayerGet(f.ctx, &pb.PlayerGetRequest{Id: id})

			assert.Nil(t, response)
			assert.Equal(t, status.Error(codes.NotFound, repository.ErrPlayerNotExists.Error()).Error(), err.Error())
		})

		t.Run("error", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			expectedError := errors.New("some error")
			var id uint64 = 10
			f.repository.EXPECT().Get(f.ctx, id).Times(1).Return(nil, expectedError)

			response, err := f.service.PlayerGet(f.ctx, &pb.PlayerGetRequest{Id: id})

			assert.Nil(t, response)
			assert.Equal(t, "rpc error: code = Internal desc = some error", err.Error())
		})
	})
}

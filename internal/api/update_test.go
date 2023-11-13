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
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestUpdate(t *testing.T) {
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

		f.repository.EXPECT().Update(f.ctx, player).Times(1).Return(nil)

		response, err := f.service.PlayerUpdate(f.ctx, &pb.PlayerUpdateRequest{
			Id:      player.Id,
			Name:    player.Name,
			Club:    player.Club,
			Games:   uint64(player.Games),
			Goals:   uint64(player.Goals),
			Assists: uint64(player.Assists),
		})

		require.NoError(t, err)
		assert.Equal(t, &emptypb.Empty{}, response)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("player not found", func(t *testing.T) {
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

			f.repository.EXPECT().Update(f.ctx, player).Times(1).Return(repository.ErrPlayerNotExists)

			response, err := f.service.PlayerUpdate(f.ctx, &pb.PlayerUpdateRequest{
				Id:      player.Id,
				Name:    player.Name,
				Club:    player.Club,
				Games:   uint64(player.Games),
				Goals:   uint64(player.Goals),
				Assists: uint64(player.Assists),
			})

			assert.Nil(t, response)
			assert.Equal(t, status.Error(codes.NotFound, repository.ErrPlayerNotExists.Error()).Error(), err.Error())
		})

		t.Run("error", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			expectedError := errors.New("some error")
			player := &model.Player{
				Id:      10,
				Name:    "Messi",
				Club:    "PSG",
				Games:   546,
				Goals:   480,
				Assists: 197,
			}

			f.repository.EXPECT().Update(f.ctx, player).Times(1).Return(expectedError)

			response, err := f.service.PlayerUpdate(f.ctx, &pb.PlayerUpdateRequest{
				Id:      player.Id,
				Name:    player.Name,
				Club:    player.Club,
				Games:   uint64(player.Games),
				Goals:   uint64(player.Goals),
				Assists: uint64(player.Assists),
			})

			assert.Nil(t, response)
			assert.Equal(t, "rpc error: code = Internal desc = some error", err.Error())
		})
	})
}

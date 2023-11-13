package api

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		var id uint64 = 10

		f.repository.EXPECT().Delete(f.ctx, id).Times(1).Return(nil)

		response, err := f.service.PlayerDelete(f.ctx, &pb.PlayerDeleteRequest{
			Id: id,
		})

		require.NoError(t, err)
		assert.Equal(t, &emptypb.Empty{}, response)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("player not found", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			var id uint64 = 10

			f.repository.EXPECT().Delete(f.ctx, id).Times(1).Return(repository.ErrPlayerNotExists)

			response, err := f.service.PlayerDelete(f.ctx, &pb.PlayerDeleteRequest{
				Id: id,
			})

			assert.Nil(t, response)
			assert.Equal(t, status.Error(codes.NotFound, repository.ErrPlayerNotExists.Error()).Error(), err.Error())
		})

		t.Run("error", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			var id uint64 = 10
			expectedError := errors.New("some error")

			f.repository.EXPECT().Delete(f.ctx, id).Times(1).Return(expectedError)

			response, err := f.service.PlayerDelete(f.ctx, &pb.PlayerDeleteRequest{
				Id: id,
			})

			assert.Nil(t, response)
			assert.Equal(t, "rpc error: code = Internal desc = some error", err.Error())
		})
	})
}

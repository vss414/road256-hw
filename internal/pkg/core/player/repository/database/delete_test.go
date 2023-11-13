package database

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := "DELETE FROM players WHERE id=$1"

		var id uint64 = 100
		f.pool.EXPECT().Exec(gomock.Any(), query, id).Times(1).Return(pgconn.CommandTag("1"), nil)

		err := f.repository.Delete(context.Background(), id)

		require.NoError(t, err)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("player not found", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			query := "DELETE FROM players WHERE id=$1"

			var id uint64 = 100
			f.pool.EXPECT().Exec(gomock.Any(), query, id).Times(1).Return(pgconn.CommandTag("0"), nil)

			err := f.repository.Delete(context.Background(), id)

			assert.Equal(
				t,
				errors.Wrapf(repository.ErrPlayerNotExists, "player id: [%d]", id).Error(),
				err.Error(),
			)
		})

		t.Run("error", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			expectedError := errors.New("some error")
			query := "DELETE FROM players WHERE id=$1"
			var id uint64 = 100

			f.pool.EXPECT().Exec(gomock.Any(), query, id).Times(1).Return(nil, expectedError)

			err := f.repository.Delete(context.Background(), id)

			assert.ErrorIs(t, expectedError, err)
		})
	})
}

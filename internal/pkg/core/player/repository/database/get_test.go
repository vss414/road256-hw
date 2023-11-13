package database

import (
	"context"
	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := "SELECT id, name, club, games, goals, assists FROM players WHERE id = $1"

		player := model.Player{
			Id:      1,
			Name:    "Messi",
			Club:    "PSG",
			Games:   546,
			Goals:   480,
			Assists: 197,
		}

		columns := []string{"id", "name", "club", "games", "goals", "assists"}
		pgxRows := pgxpoolmock.NewRows(columns).AddRow(
			player.Id,
			player.Name,
			player.Club,
			player.Games,
			player.Goals,
			player.Assists,
		).ToPgxRows()

		var id uint64 = 100
		f.pool.EXPECT().Query(gomock.Any(), query, id).Times(1).Return(pgxRows, nil)

		result, err := f.repository.Get(context.Background(), id)

		require.NoError(t, err)
		assert.Equal(t, player.Id, result.Id)
		assert.Equal(t, player.Name, result.Name)
		assert.Equal(t, player.Club, result.Club)
		assert.Equal(t, player.Games, result.Games)
		assert.Equal(t, player.Goals, result.Goals)
		assert.Equal(t, player.Assists, result.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("player not found", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			query := "SELECT id, name, club, games, goals, assists FROM players WHERE id = $1"

			var id uint64 = 100
			f.pool.EXPECT().
				Query(gomock.Any(), query, id).
				Times(1).
				Return(pgxpoolmock.NewRows([]string{}).ToPgxRows(), nil)

			player, err := f.repository.Get(context.Background(), id)

			assert.Nil(t, player)
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

			query := "SELECT id, name, club, games, goals, assists FROM players WHERE id = $1"

			var id uint64 = 100
			f.pool.EXPECT().Query(gomock.Any(), query, id).Times(1).Return(nil, expectedError)

			player, err := f.repository.Get(context.Background(), id)

			assert.Nil(t, player)
			assert.Equal(t, "scany: query one result row: some error", err.Error())
		})
	})
}

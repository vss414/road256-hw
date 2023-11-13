package database

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := "UPDATE players SET name = $1, club = $2, games = $3, goals = $4, assists = $5 WHERE id = $6"
		player := model.Player{
			Id:      10,
			Name:    "Messi",
			Club:    "PSG",
			Games:   546,
			Goals:   480,
			Assists: 197,
		}

		f.pool.EXPECT().
			Exec(gomock.Any(), query, player.Name, player.Club, player.Games, player.Goals, player.Assists, player.Id).
			Times(1).
			Return(pgconn.CommandTag("1"), nil)

		err := f.repository.Update(context.Background(), &player)

		require.NoError(t, err)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("player not found", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			query := "UPDATE players SET name = $1, club = $2, games = $3, goals = $4, assists = $5 WHERE id = $6"
			player := model.Player{
				Id:      10,
				Name:    "Messi",
				Club:    "PSG",
				Games:   546,
				Goals:   480,
				Assists: 197,
			}

			f.pool.EXPECT().
				Exec(
					gomock.Any(),
					query,
					player.Name,
					player.Club,
					player.Games,
					player.Goals,
					player.Assists,
					player.Id,
				).
				Times(1).
				Return(pgconn.CommandTag("0"), nil)

			err := f.repository.Update(context.Background(), &player)

			assert.Equal(
				t,
				errors.Wrapf(repository.ErrPlayerNotExists, "player id: [%d]", player.Id).Error(),
				err.Error(),
			)
		})

		t.Run("error", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			expectedError := errors.New("some error")
			query := "UPDATE players SET name = $1, club = $2, games = $3, goals = $4, assists = $5 WHERE id = $6"
			player := model.Player{
				Id:      10,
				Name:    "Messi",
				Club:    "PSG",
				Games:   546,
				Goals:   480,
				Assists: 197,
			}

			f.pool.EXPECT().
				Exec(
					gomock.Any(),
					query,
					player.Name,
					player.Club,
					player.Games,
					player.Goals,
					player.Assists,
					player.Id,
				).
				Times(1).
				Return(nil, expectedError)

			err := f.repository.Update(context.Background(), &player)

			assert.ErrorIs(t, expectedError, err)
		})
	})
}

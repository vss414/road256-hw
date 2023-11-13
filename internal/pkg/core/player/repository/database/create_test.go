package database

import (
	"context"
	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := "INSERT INTO public.players (name, club, games, goals, assists) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		player := model.Player{
			Name:    "Messi",
			Club:    "PSG",
			Games:   546,
			Goals:   480,
			Assists: 197,
		}
		var id uint64 = 10

		f.pool.EXPECT().
			QueryRow(gomock.Any(), query, player.Name, player.Club, player.Games, player.Goals, player.Assists).
			Times(1).
			Return(pgxpoolmock.NewRow(id))

		result, err := f.repository.Create(context.Background(), &player)

		require.NoError(t, err)
		assert.Equal(t, id, result.Id)
		assert.Equal(t, player.Name, result.Name)
		assert.Equal(t, player.Club, result.Club)
		assert.Equal(t, player.Games, result.Games)
		assert.Equal(t, player.Goals, result.Goals)
		assert.Equal(t, player.Assists, result.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := "INSERT INTO public.players (name, club, games, goals, assists) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		player := model.Player{
			Name:    "Messi",
			Club:    "PSG",
			Games:   546,
			Goals:   480,
			Assists: 197,
		}
		expectedError := errors.New("some error")

		f.pool.EXPECT().
			QueryRow(gomock.Any(), query, player.Name, player.Club, player.Games, player.Goals, player.Assists).
			Times(1).
			Return(pgxpoolmock.NewRow(uint64(0)).WithError(expectedError))

		result, err := f.repository.Create(context.Background(), &player)

		assert.Nil(t, result)
		assert.Error(t, err)
	})
}

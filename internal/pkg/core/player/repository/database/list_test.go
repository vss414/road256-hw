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

func TestList(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := "SELECT id, name, club, games, goals, assists FROM players ORDER BY ID ASC LIMIT 3 OFFSET 1"
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

		f.pool.EXPECT().Query(gomock.Any(), query).Times(1).Return(pgxRows, nil)

		result, err := f.repository.List(context.Background(), repository.Pagination{
			Limit:     3,
			Offset:    1,
			Order:     "ID",
			Direction: "ASC",
		})

		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, player.Id, result[0].Id)
		assert.Equal(t, player.Name, result[0].Name)
		assert.Equal(t, player.Club, result[0].Club)
		assert.Equal(t, player.Games, result[0].Games)
		assert.Equal(t, player.Goals, result[0].Goals)
		assert.Equal(t, player.Assists, result[0].Assists)
	})

	t.Run("negative", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		expectedError := errors.New("some error")

		query := "SELECT id, name, club, games, goals, assists FROM players ORDER BY ID ASC LIMIT 3 OFFSET 1"

		f.pool.EXPECT().Query(gomock.Any(), query).Times(1).Return(nil, expectedError)

		result, err := f.repository.List(context.Background(), repository.Pagination{
			Limit:     3,
			Offset:    1,
			Order:     "ID",
			Direction: "ASC",
		})

		assert.Nil(t, result)
		assert.Equal(t, "scany: query multiple result rows: some error", err.Error())
	})
}

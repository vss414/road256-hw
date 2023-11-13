//go:build integration
// +build integration

package tests

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"math"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		r := player.New(Pool)

		p := &model.Player{}
		err := faker.FakeData(&p)
		require.NoError(t, err)

		result, err := r.Create(context.Background(), p)
		require.NoError(t, err)

		assert.Greater(t, result.Id, uint64(0))
		assert.Equal(t, p.Name, result.Name)
		assert.Equal(t, p.Club, result.Club)
		assert.Equal(t, p.Games, result.Games)
		assert.Equal(t, p.Goals, result.Goals)
		assert.Equal(t, p.Assists, result.Assists)
	})
}

func TestGet(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		r := player.New(Pool)

		createdPlayer := createInDb(t, r)

		p, err := r.Get(context.Background(), createdPlayer.Id)
		require.NoError(t, err)
		assert.Equal(t, createdPlayer.Id, p.Id)
		assert.Equal(t, createdPlayer.Name, p.Name)
		assert.Equal(t, createdPlayer.Club, p.Club)
		assert.Equal(t, createdPlayer.Games, p.Games)
		assert.Equal(t, createdPlayer.Goals, p.Goals)
		assert.Equal(t, createdPlayer.Assists, p.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		r := player.New(Pool)

		id := uint64(math.MaxInt32)

		p, err := r.Get(context.Background(), id)

		assert.Nil(t, p)
		expectedError := fmt.Sprintf("player id: [%d]: player does not exist", id)
		assert.Equal(t, expectedError, err.Error())
	})
}

func TestList(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("ASC direction", func(t *testing.T) {
			r := player.New(Pool)

			createInDb(t, r)
			createInDb(t, r)
			createInDb(t, r)
			createInDb(t, r)

			players, err := r.List(context.Background(), 3, 1, pb.Order_ORDER_ID.String(), pb.Direction_DIRECTION_ASC.String())

			require.NoError(t, err)
			assert.Len(t, players, 3)

			for i, p := range players {
				if i < len(players)-1 {
					assert.Greater(t, players[i+1].Id, p.Id)
				}
			}
		})

		t.Run("DESC direction", func(t *testing.T) {
			r := player.New(Pool)

			createInDb(t, r)
			createInDb(t, r)
			createInDb(t, r)
			createInDb(t, r)

			players, err := r.List(
				context.Background(),
				3,
				1,
				pb.Order_ORDER_GAMES.String(),
				pb.Direction_DIRECTION_DESC.String(),
			)

			require.NoError(t, err)
			assert.Len(t, players, 3)

			for i, p := range players {
				if i < len(players)-1 {
					assert.GreaterOrEqual(t, p.Games, players[i+1].Games)
				}
			}
		})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		r := player.New(Pool)

		createdPlayer := createInDb(t, r)

		updatedPlayer := &model.Player{}
		err := faker.FakeData(&updatedPlayer)
		require.NoError(t, err)
		updatedPlayer.Id = createdPlayer.Id

		err = r.Update(context.Background(), updatedPlayer)
		require.NoError(t, err)

		p, err := r.Get(context.Background(), createdPlayer.Id)
		require.NoError(t, err)
		assert.Equal(t, updatedPlayer.Id, p.Id)
		assert.Equal(t, updatedPlayer.Name, p.Name)
		assert.Equal(t, updatedPlayer.Club, p.Club)
		assert.Equal(t, updatedPlayer.Games, p.Games)
		assert.Equal(t, updatedPlayer.Goals, p.Goals)
		assert.Equal(t, updatedPlayer.Assists, p.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("wrong name", func(t *testing.T) {
			r := player.New(Pool)

			createdPlayer := createInDb(t, r)
			updatedPlayer := &model.Player{}
			updatedPlayer.Id = createdPlayer.Id

			err := r.Update(context.Background(), updatedPlayer)

			expectedError := "field: [name] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
		})

		t.Run("wrong club", func(t *testing.T) {
			r := player.New(Pool)

			createdPlayer := createInDb(t, r)
			updatedPlayer := &model.Player{}
			updatedPlayer.Id = createdPlayer.Id
			err := faker.FakeData(&updatedPlayer.Name)
			require.NoError(t, err)

			err = r.Update(context.Background(), updatedPlayer)

			expectedError := "field: [club] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
		})
	})
}

func TestDelete(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		r := player.New(Pool)

		createdPlayer := createInDb(t, r)
		err := r.Delete(context.Background(), createdPlayer.Id)
		require.NoError(t, err)
	})

	t.Run("negative", func(t *testing.T) {
		r := player.New(Pool)
		id := uint64(math.MaxInt32)

		err := r.Delete(context.Background(), id)

		expectedError := fmt.Sprintf("player id: [%d]: player does not exist", id)
		assert.Equal(t, expectedError, err.Error())
	})
}

func createInDb(t *testing.T, r player.IPlayer) *model.Player {
	p := &model.Player{}
	err := faker.FakeData(&p)
	require.NoError(t, err)

	result, err := r.Create(context.Background(), p)
	require.NoError(t, err)

	return result
}

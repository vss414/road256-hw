package facade

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	"testing"
)

func TestFacade_Get(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("get from database", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			var id uint64 = 10
			player := &model.Player{
				Id:      1,
				Name:    "Messi",
				Club:    "PSG",
				Games:   546,
				Goals:   480,
				Assists: 197,
			}

			f.local.EXPECT().Get(f.ctx, id).Times(1).Return(nil, repositoryPkg.ErrPlayerNotExists)
			f.database.EXPECT().Get(f.ctx, id).Times(1).Return(player, nil)
			f.local.EXPECT().Create(f.ctx, player).Times(1).Return(player, nil)

			result, err := f.facade.Get(context.Background(), 10)
			require.NoError(t, err)

			assert.Equal(t, player.Id, result.Id)
			assert.Equal(t, player.Name, result.Name)
			assert.Equal(t, player.Club, result.Club)
			assert.Equal(t, player.Games, result.Games)
			assert.Equal(t, player.Goals, result.Goals)
			assert.Equal(t, player.Assists, result.Assists)
		})

		t.Run("get from local", func(t *testing.T) {
			f := setUp(t)
			defer f.tearDown()

			var id uint64 = 10
			player := &model.Player{
				Id:      1,
				Name:    "Messi",
				Club:    "PSG",
				Games:   546,
				Goals:   480,
				Assists: 197,
			}

			f.local.EXPECT().Get(f.ctx, id).Times(1).Return(player, nil)
			f.database.EXPECT().Get(f.ctx, id).Times(0)
			f.local.EXPECT().Create(f.ctx, player).Times(0)

			result, err := f.facade.Get(context.Background(), 10)
			require.NoError(t, err)

			assert.Equal(t, player.Id, result.Id)
			assert.Equal(t, player.Name, result.Name)
			assert.Equal(t, player.Club, result.Club)
			assert.Equal(t, player.Games, result.Games)
			assert.Equal(t, player.Goals, result.Goals)
			assert.Equal(t, player.Assists, result.Assists)
		})
	})
}

func TestFacade_List(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		player := &model.Player{
			Id:      1,
			Name:    "Messi",
			Club:    "PSG",
			Games:   546,
			Goals:   480,
			Assists: 197,
		}
		pagination := repositoryPkg.Pagination{
			Limit:     3,
			Offset:    1,
			Order:     "ID",
			Direction: "ASC",
		}

		f.database.EXPECT().List(f.ctx, pagination).Times(1).Return([]*model.Player{player, player, player}, nil)
		f.local.EXPECT().Create(f.ctx, player).Times(3).Return(player, nil)

		result, err := f.facade.List(context.Background(), pagination)
		require.NoError(t, err)

		assert.Equal(t, player.Id, result[0].Id)
		assert.Equal(t, player.Name, result[0].Name)
		assert.Equal(t, player.Club, result[0].Club)
		assert.Equal(t, player.Games, result[0].Games)
		assert.Equal(t, player.Goals, result[0].Goals)
		assert.Equal(t, player.Assists, result[0].Assists)
	})
}

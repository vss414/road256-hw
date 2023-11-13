package database

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func (d database) Get(ctx context.Context, id uint64) (*model.Player, error) {
	query := "SELECT id, name, club, games, goals, assists FROM players WHERE id = $1"

	var player model.Player
	if err := pgxscan.Get(ctx, d.pool, &player, query, id); err != nil {
		if pgxscan.NotFound(err) {
			return nil, errors.Wrapf(repositoryPkg.ErrPlayerNotExists, "player id: [%d]", id)
		}

		return nil, err
	}

	return &player, nil
}

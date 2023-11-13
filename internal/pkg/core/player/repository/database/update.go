package database

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func (d database) Update(ctx context.Context, p *model.Player) error {
	query := "UPDATE players SET name = $1, club = $2, games = $3, goals = $4, assists = $5 WHERE id = $6"

	row, err := d.pool.Exec(ctx, query, p.Name, p.Club, p.Games, p.Goals, p.Assists, p.Id)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return errors.Wrapf(repositoryPkg.ErrPlayerNotExists, "player id: [%d]", p.Id)
	}

	return nil
}

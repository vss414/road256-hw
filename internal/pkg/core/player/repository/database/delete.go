package database

import (
	"context"
	"github.com/pkg/errors"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func (d database) Delete(ctx context.Context, id uint64) error {
	query := "DELETE FROM players WHERE id=$1"

	row, err := d.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return errors.Wrapf(repositoryPkg.ErrPlayerNotExists, "player id: [%d]", id)
	}

	return nil
}

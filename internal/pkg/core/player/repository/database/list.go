package database

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func (d database) List(ctx context.Context, p repository.Pagination) ([]*model.Player, error) {
	query, args, err := squirrel.Select("id, name, club, games, goals, assists").
		From("players").
		Limit(p.Limit).
		Offset(p.Offset).
		OrderBy(fmt.Sprintf("%s %s", p.Order, p.Direction)).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("prepare list query: %w", err)
	}

	var players []*model.Player
	if err := pgxscan.Select(ctx, d.pool, &players, query, args...); err != nil {
		return nil, err
	}

	return players, nil
}

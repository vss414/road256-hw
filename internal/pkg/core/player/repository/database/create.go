package database

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
)

func (d database) Create(ctx context.Context, p *model.Player) (*model.Player, error) {
	query := "INSERT INTO public.players (name, club, games, goals, assists) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	row := d.pool.QueryRow(ctx, query, p.Name, p.Club, p.Games, p.Goals, p.Assists)
	if err := row.Scan(&p.Id); err != nil {
		return nil, fmt.Errorf("scan insert: %w", err)
	}

	return p, nil
}

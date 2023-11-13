package consumer

import (
	"context"
	playerPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player"
)

type IHandler interface {
	Create(ctx context.Context, in []byte) error
	Update(ctx context.Context, in []byte) error
	Delete(ctx context.Context, in []byte) error
}

type handler struct {
	player playerPkg.IPlayer
}

func New(player playerPkg.IPlayer) IHandler {
	return &handler{
		player: player,
	}
}

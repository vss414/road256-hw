package consumer

import (
	"context"
	"encoding/json"
	playerModelPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
)

func (h *handler) Update(ctx context.Context, in []byte) error {
	var request pb.PlayerUpdateRequest
	if err := json.Unmarshal(in, &request); err != nil {
		return err
	}

	p := &playerModelPkg.Player{
		Id:      request.GetId(),
		Name:    request.GetName(),
		Club:    request.GetClub(),
		Games:   uint(request.GetGames()),
		Goals:   uint(request.GetGoals()),
		Assists: uint(request.GetAssists()),
	}

	return h.player.Update(ctx, p)
}

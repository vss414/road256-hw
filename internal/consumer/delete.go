package consumer

import (
	"context"
	"encoding/json"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
)

func (h *handler) Delete(ctx context.Context, in []byte) error {
	var request pb.PlayerDeleteRequest
	if err := json.Unmarshal(in, &request); err != nil {
		return err
	}

	return h.player.Delete(ctx, request.Id)
}

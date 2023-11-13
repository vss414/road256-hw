package local

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func (c *repository) Update(ctx context.Context, player *model.Player) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	if _, ok := c.data[player.Id]; !ok {
		return errors.Wrapf(repositoryPkg.ErrPlayerNotExists, "player id: [%d]", player.Id)
	}

	c.data[player.Id] = player
	return nil
}

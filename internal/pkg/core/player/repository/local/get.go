package local

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func (c *repository) Get(ctx context.Context, id uint64) (*model.Player, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
	}()

	if p, ok := c.data[id]; ok {
		return p, nil
	}

	return nil, errors.Wrapf(repositoryPkg.ErrPlayerNotExists, "player id: [%d]", id)
}

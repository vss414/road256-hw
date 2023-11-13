package local

import (
	"context"
	"github.com/pkg/errors"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func (c *repository) Delete(ctx context.Context, id uint64) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	if _, ok := c.data[id]; !ok {
		return errors.Wrapf(repositoryPkg.ErrPlayerNotExists, "player name: [%d]", id)
	}

	delete(c.data, id)
	return nil
}

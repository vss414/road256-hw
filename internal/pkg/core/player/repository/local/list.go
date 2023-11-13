package local

import (
	"context"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func (c *repository) List(ctx context.Context, p repositoryPkg.Pagination) ([]*model.Player, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
	}()

	keys := make([]uint64, 0, p.Limit)

	var i uint64
	for k := range c.data {
		if i >= p.Offset {
			keys = append(keys, k)
		}
		i++
		if i >= p.Offset+p.Limit {
			break
		}
	}

	players := make([]*model.Player, 0, len(keys))
	for _, k := range keys {
		players = append(players, c.data[k])
	}
	return players, nil
}

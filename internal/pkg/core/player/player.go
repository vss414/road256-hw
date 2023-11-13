//go:generate mockgen -source ./player.go -destination=./mocks/player.go -package=mock_player

package player

import (
	"context"
	"fmt"
	"github.com/chrisyxlee/pgxpoolmock"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository/facade"
	"time"
)

type IPlayer interface {
	List(ctx context.Context, limit, page uint64, order, direction string) ([]*model.Player, error)
	Get(ctx context.Context, id uint64) (*model.Player, error)
	Create(ctx context.Context, p *model.Player) (*model.Player, error)
	Update(ctx context.Context, p *model.Player) error
	Delete(ctx context.Context, id uint64) error
}

func New(pool pgxpoolmock.PgxPool) IPlayer {
	return &core{
		facade: facade.New(pool),
	}
}

type core struct {
	facade facade.IFacade
}

func (c *core) List(ctx context.Context, limit, page uint64, order, direction string) ([]*model.Player, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	ch := make(chan []*model.Player, 1)
	errCh := make(chan error, 1)

	go func() {
		pagination, err := repositoryPkg.NewPagination(limit, page, order, direction)
		if err != nil {
			errCh <- err
		}

		p, err := c.facade.List(ctx, pagination)
		if err != nil {
			errCh <- err
		}
		ch <- p
	}()

	select {
	case p := <-ch:
		return p, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return nil, ctx.Err()
	}
}

func (c *core) Get(ctx context.Context, id uint64) (*model.Player, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	ch := make(chan *model.Player, 1)
	errCh := make(chan error, 1)

	go func() {
		p, err := c.facade.Get(ctx, id)
		if err != nil {
			errCh <- err
		}
		ch <- p
	}()

	select {
	case p := <-ch:
		return p, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return nil, ctx.Err()
	}
}

func (c *core) Create(ctx context.Context, p *model.Player) (*model.Player, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	ch := make(chan *model.Player, 1)
	errCh := make(chan error, 1)

	go func() {
		if err := p.Validate(); err != nil {
			errCh <- err
		}

		player, err := c.facade.Create(ctx, p)
		if err != nil {
			errCh <- err
		}
		ch <- player
	}()

	select {
	case player := <-ch:
		return player, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return nil, ctx.Err()
	}
}

func (c *core) Update(ctx context.Context, p *model.Player) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	errCh := make(chan error, 1)

	go func() {
		if err := p.Validate(); err != nil {
			errCh <- err
		}

		if err := c.facade.Update(ctx, p); err != nil {
			errCh <- err
		}
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return ctx.Err()
	}
}

func (c *core) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	errCh := make(chan error, 1)

	go func() {
		if err := c.facade.Delete(ctx, id); err != nil {
			errCh <- err
		}
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return ctx.Err()
	}
}

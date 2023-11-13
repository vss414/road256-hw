package facade

import (
	"context"
	"fmt"
	"github.com/chrisyxlee/pgxpoolmock"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository/database"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository/local"
)

type IFacade interface {
	List(ctx context.Context, p repositoryPkg.Pagination) ([]*model.Player, error)
	Get(ctx context.Context, id uint64) (*model.Player, error)
	Create(ctx context.Context, p *model.Player) (*model.Player, error)
	Update(ctx context.Context, p *model.Player) error
	Delete(ctx context.Context, id uint64) error
}

func New(pool pgxpoolmock.PgxPool) IFacade {
	return &facade{
		local:    local.New(),
		database: database.New(pool),
	}
}

type facade struct {
	local    repositoryPkg.IRepository
	database repositoryPkg.IRepository
}

func (s *facade) List(ctx context.Context, p repositoryPkg.Pagination) ([]*model.Player, error) {
	players, err := s.database.List(ctx, p)

	if err == nil {
		for _, p := range players {
			if _, err := s.local.Create(ctx, p); err != nil {
				fmt.Printf("failed to create player <%d> in local cache: %s\n", p.Id, err)
			}
		}
	}

	return players, err

}

func (s *facade) Get(ctx context.Context, id uint64) (*model.Player, error) {
	if p, err := s.local.Get(ctx, id); err == nil {
		fmt.Printf("get player <%d> from cache\n", id)
		return p, nil
	}

	p, err := s.database.Get(ctx, id)
	if err == nil {
		if _, err := s.local.Create(ctx, p); err != nil {
			fmt.Printf("failed to create player <%d> in local cache: %s\n", p.Id, err)
		}
	}

	return p, err
}

func (s *facade) Create(ctx context.Context, p *model.Player) (*model.Player, error) {
	player, err := s.database.Create(ctx, p)

	if err == nil {
		if _, err := s.local.Create(ctx, player); err != nil {
			fmt.Printf("failed to create player <%d> in local cache: %s\n", player.Id, err)
		}
	}

	return player, err
}

func (s *facade) Update(ctx context.Context, p *model.Player) error {
	err := s.database.Update(ctx, p)
	if err == nil {
		if err := s.local.Update(ctx, p); err != nil {
			fmt.Printf("failed to update player <%d> in local cache: %s\n", p.Id, err)
		}
	}

	return err
}

func (s *facade) Delete(ctx context.Context, id uint64) error {
	err := s.database.Delete(ctx, id)

	if err == nil {
		if err := s.local.Delete(ctx, id); err != nil {
			fmt.Printf("failed to delete player <%d> from local cache: %s\n", id, err)
		}
	}

	return err
}

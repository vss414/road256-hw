//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository

package repository

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
)

var (
	ErrListParameter   = errors.New("wrong list parameter")
	ErrPlayerExists    = errors.New("player exists")
	ErrPlayerNotExists = errors.New("player does not exist")
)

type IRepository interface {
	List(ctx context.Context, p Pagination) ([]*model.Player, error)
	Get(ctx context.Context, id uint64) (*model.Player, error)
	Create(ctx context.Context, p *model.Player) (*model.Player, error)
	Update(ctx context.Context, p *model.Player) error
	Delete(ctx context.Context, id uint64) error
}

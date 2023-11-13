package database

import (
	"github.com/chrisyxlee/pgxpoolmock"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
)

func New(pool pgxpoolmock.PgxPool) repositoryPkg.IRepository {
	return &database{pool}
}

type database struct {
	pool pgxpoolmock.PgxPool
}

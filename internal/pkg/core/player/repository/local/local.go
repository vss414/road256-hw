package local

import (
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
	repositoryPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	"sync"
)

const poolSize = 10

func New() repositoryPkg.IRepository {
	return &repository{
		mu:     sync.RWMutex{},
		data:   make(map[uint64]*model.Player),
		poolCh: make(chan struct{}, poolSize),
	}
}

type repository struct {
	mu     sync.RWMutex
	data   map[uint64]*model.Player
	poolCh chan struct{}
}

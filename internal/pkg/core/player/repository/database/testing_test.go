package database

import (
	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository"
	"testing"
)

type repositoryFixtures struct {
	controller *gomock.Controller
	repository repository.IRepository
	pool       *pgxpoolmock.MockPgxIface
}

func setUp(t *testing.T) repositoryFixtures {
	controller := gomock.NewController(t)
	pool := pgxpoolmock.NewMockPgxIface(controller)

	return repositoryFixtures{
		controller: controller,
		repository: New(pool),
		pool:       pool,
	}
}

func (f *repositoryFixtures) tearDown() {
	f.controller.Finish()
}

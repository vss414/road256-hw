package facade

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_repository "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/repository/mocks"
	"testing"
)

type facadeFixtures struct {
	ctx        context.Context
	controller *gomock.Controller
	facade     IFacade
	local      *mock_repository.MockIRepository
	database   *mock_repository.MockIRepository
}

func setUp(t *testing.T) *facadeFixtures {
	controller := gomock.NewController(t)

	local := mock_repository.NewMockIRepository(controller)
	database := mock_repository.NewMockIRepository(controller)

	return &facadeFixtures{
		ctx:        context.Background(),
		controller: controller,
		local:      local,
		database:   database,
		facade: &facade{
			local:    local,
			database: database,
		},
	}
}

func (f *facadeFixtures) tearDown() {
	f.controller.Finish()
}

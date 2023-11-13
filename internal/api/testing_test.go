package api

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	mock_player "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/mocks"
	"gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"testing"
)

type serviceFixtures struct {
	ctx        context.Context
	controller *gomock.Controller
	repository *mock_player.MockIPlayer
	service    api.AdminServer
}

func setUp(t *testing.T) serviceFixtures {
	controller := gomock.NewController(t)
	player := mock_player.NewMockIPlayer(controller)
	logger, _ := test.NewNullLogger()

	return serviceFixtures{
		ctx:        context.Background(),
		controller: controller,
		repository: player,
		service: &implementation{
			player: player,
			log:    logger,
		},
	}
}

func (f *serviceFixtures) tearDown() {
	f.controller.Finish()
}

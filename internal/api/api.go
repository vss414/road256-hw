package api

import (
	"github.com/sirupsen/logrus"
	loggerPkg "gitlab.ozon.dev/vss414/hw-1/internal/logger"
	playerPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"log"
)

func New(player playerPkg.IPlayer) pb.AdminServer {
	logger, err := loggerPkg.New("repository.log")
	if err != nil {
		log.Fatal(err)
	}

	return &implementation{
		player: player,
		log:    logger,
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	player playerPkg.IPlayer
	log    *logrus.Logger
}

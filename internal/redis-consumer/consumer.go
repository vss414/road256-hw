package redis_consumer

import (
	"github.com/sirupsen/logrus"
	loggerPkg "gitlab.ozon.dev/vss414/hw-1/internal/logger"
	playerPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player"
	"gitlab.ozon.dev/vss414/hw-1/internal/redis"
	"log"
)

type IRedisConsumer interface {
	Handle()
	handleGet()
	handleList()
}

type consumer struct {
	player playerPkg.IPlayer
	redis  redis.IRedis
	logger *logrus.Logger
}

func New(player playerPkg.IPlayer) IRedisConsumer {
	logger, err := loggerPkg.New("repository.log")
	if err != nil {
		log.Fatal(err)
	}

	return &consumer{
		player: player,
		redis:  redis.New(),
		logger: logger,
	}
}

func (c *consumer) Handle() {
	go c.handleGet()
	go c.handleList()
}

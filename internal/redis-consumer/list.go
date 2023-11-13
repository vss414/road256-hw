package redis_consumer

import (
	"context"
	"encoding/json"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
)

const (
	PlayerList         = "player_list"
	PlayerListResponse = "player_list_response"
	PlayerListError    = "player_list_error"
)

func (c *consumer) handleList() {
	method := "handle-player-list"
	pubsub := c.redis.Subscribe(PlayerList)
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		c.logger.Infof("%s: message: %s", method, msg.String())

		var request *pb.PlayerListRequest
		if err := json.Unmarshal([]byte(msg.Payload), &request); err != nil {
			c.logger.Errorf("%s: message: %s; unmarshal error: %s", method, msg.Payload, err)
			if err := c.redis.Publish(PlayerListError, err.Error()); err != nil {
				c.logger.Errorf("%s: publish error: %s", method, err)
			}
			continue
		}

		players, err := c.player.List(
			context.Background(),
			request.GetLimit(),
			request.GetPage(),
			request.GetOrder().String(),
			request.GetDirection().String(),
		)
		if err != nil {
			c.logger.Errorf("%s: player list error: %s", method, err)
			if err := c.redis.Publish(PlayerListError, err.Error()); err != nil {
				c.logger.Errorf("%s: publish error: %s", method, err)
			}
			continue
		}

		redisMessage, err := json.Marshal(players)
		if err != nil {
			c.logger.Errorf("%s: marshal error: %s", method, err)
			if err := c.redis.Publish(PlayerListError, err.Error()); err != nil {
				c.logger.Errorf("%s: publish error: %s", method, err)
			}
			continue
		}

		if err := c.redis.Publish(PlayerListResponse, redisMessage); err != nil {
			c.logger.Errorf("%s: publish response error: %s", method, err)
			if err := c.redis.Publish(PlayerListError, err.Error()); err != nil {
				c.logger.Errorf("%s: publish error: %s", method, err)
			}
			continue
		}

		c.logger.Infof("%s: %s Done!", method, msg.String())
	}
}

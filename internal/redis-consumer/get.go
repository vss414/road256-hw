package redis_consumer

import (
	"context"
	"encoding/json"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
)

const (
	PlayerGet         = "player_get"
	PlayerGetResponse = "player_get_response"
	PlayerGetError    = "player_get_error"
)

func (c *consumer) handleGet() {
	method := "handle-player-get"
	pubsub := c.redis.Subscribe(PlayerGet)
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		c.logger.Infof("%s: message: %s", method, msg.String())

		var request *pb.PlayerGetRequest
		if err := json.Unmarshal([]byte(msg.Payload), &request); err != nil {
			c.logger.Errorf("%s: message: %s; unmarshal error: %s", method, msg.Payload, err)
			if err := c.redis.Publish(PlayerGetError, err.Error()); err != nil {
				c.logger.Errorf("%s: publish error: %s", method, err)
			}
			continue
		}

		player, err := c.player.Get(context.Background(), request.Id)
		if err != nil {
			c.logger.Errorf("%s: player get error: %s", method, err)
			if err := c.redis.Publish(PlayerGetError, err.Error()); err != nil {
				c.logger.Errorf("%s: publish error: %s", method, err)
			}
			continue
		}

		redisMessage, err := json.Marshal(player)
		if err != nil {
			c.logger.Errorf("%s: marshal error: %s", method, err)
			if err := c.redis.Publish(PlayerGetError, err.Error()); err != nil {
				c.logger.Errorf("%s: publish error: %s", method, err)
			}
			continue
		}

		if err := c.redis.Publish(PlayerGetResponse, redisMessage); err != nil {
			c.logger.Errorf("%s: publish response error: %s", method, err)
			if err := c.redis.Publish(PlayerGetError, err.Error()); err != nil {
				c.logger.Errorf("%s: publish error: %s", method, err)
			}
			continue
		}

		c.logger.Infof("%s: %s Done!", method, msg.String())
	}
}

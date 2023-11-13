package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	playerPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player"
	"log"
	"regexp"
)

var PlayerEventsTopic = "player-events"

type Consumer struct {
	h IHandler
}

func (c *Consumer) handle(ctx context.Context, msg *sarama.ConsumerMessage) error {
	switch key := msg.Key; {
	case regexp.MustCompile(`^create`).MatchString(string(key)):
		if err := c.h.Create(ctx, msg.Value); err != nil {
			return errors.Wrap(err, "create player")
		}
	case regexp.MustCompile(`^update`).MatchString(string(key)):
		if err := c.h.Update(ctx, msg.Value); err != nil {
			return errors.Wrap(err, "update player")
		}
	case regexp.MustCompile(`^delete`).MatchString(string(key)):
		if err := c.h.Delete(ctx, msg.Value); err != nil {
			return errors.Wrap(err, "delete player")
		}
	}

	return nil
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case <-session.Context().Done():
			log.Print("Done")
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Print("Data channel closed")
				return nil
			}

			if err := c.handle(session.Context(), msg); err != nil {
				log.Printf("handle message: %v", err)
			}

			session.MarkMessage(msg, "")
		}
	}
}

func Consume(player playerPkg.IPlayer) {
	brokers := []string{"localhost:19091", "localhost:29091", "localhost:39091"}
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokers, "startConsuming", config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx := context.Background()

	consumer := &Consumer{
		h: New(player),
	}

	for {
		if err := client.Consume(ctx, []string{PlayerEventsTopic}, consumer); err != nil {
			log.Printf("on consume: %v", err)
		}
	}
}

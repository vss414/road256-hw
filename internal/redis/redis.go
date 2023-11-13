package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type IRedis interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, dest interface{}) error
	Delete(key string) error
	Publish(channel string, message interface{}) error
	Subscribe(channel ...string) *redis.PubSub
}

func New() IRedis {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:63790",
		DB:       0,
		Password: "",
	})

	if redisClient == nil {
		log.Fatal("failed to init redis")
	}

	return &implementation{
		redis: redisClient,
	}
}

type implementation struct {
	redis *redis.Client
}

func (i *implementation) Set(key string, value interface{}, expiration time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return i.redis.Set(key, p, expiration).Err()
}

func (i *implementation) Get(key string, dest interface{}) error {
	p, err := i.redis.Get(key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(p, dest)
}

func (i *implementation) Delete(key string) error {
	return i.redis.Del(key).Err()
}

func (i *implementation) Publish(channel string, message interface{}) error {
	return i.redis.Publish(channel, message).Err()
}

func (i *implementation) Subscribe(channel ...string) *redis.PubSub {
	return i.redis.Subscribe(channel...)
}

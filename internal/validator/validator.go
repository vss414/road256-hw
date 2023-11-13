package validator

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	loggerPkg "gitlab.ozon.dev/vss414/hw-1/internal/logger"
	"gitlab.ozon.dev/vss414/hw-1/internal/redis"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func New() pb.AdminServer {
	clientConn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewAdminClient(clientConn)

	brokers := []string{"localhost:19091", "localhost:29091", "localhost:39091"}
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	asyncProducer, err := sarama.NewAsyncProducer(brokers, cfg)
	if err != nil {
		log.Fatalf("initial async kafka: %v", err)
	}

	go func() {
		for msg := range asyncProducer.Errors() {
			log.Printf("error: %v", msg)
		}
	}()

	go func() {
		for res := range asyncProducer.Successes() {
			fmt.Printf("success: %+v\n", res)
		}
	}()

	logger, err := loggerPkg.New("validator.log")
	if err != nil {
		log.Fatal(err)
	}

	return &implementation{
		client:   client,
		producer: asyncProducer,
		log:      logger,
		redis:    redis.New(),
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	client   pb.AdminClient
	producer sarama.AsyncProducer
	log      *logrus.Logger
	redis    redis.IRedis
}

func (i *implementation) redisKey(id uint64) string {
	return fmt.Sprintf("player:%d", id)
}

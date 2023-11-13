package main

import (
	apiPkg "gitlab.ozon.dev/vss414/hw-1/internal/api"
	"gitlab.ozon.dev/vss414/hw-1/internal/consumer"
	_ "gitlab.ozon.dev/vss414/hw-1/internal/counter"
	"gitlab.ozon.dev/vss414/hw-1/internal/database"
	playerPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player"
	redis_consumer "gitlab.ozon.dev/vss414/hw-1/internal/redis-consumer"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	// Counter
	go func() {
		if err := http.ListenAndServe(":8088", nil); err != nil {
			log.Fatal(err)
		}
	}()

	pool := database.New()
	defer pool.Close()

	player := playerPkg.New(pool)

	go consumer.Consume(player)

	redisConsumer := redis_consumer.New(player)
	go redisConsumer.Handle()

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(player))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

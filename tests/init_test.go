//go:build integration
// +build integration

package tests

import (
	"github.com/chrisyxlee/pgxpoolmock"
	"gitlab.ozon.dev/vss414/hw-1/internal/database"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	Client pb.AdminClient
	Pool   pgxpoolmock.PgxPool
)

func init() {
	clientConn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	Client = pb.NewAdminClient(clientConn)

	Pool = database.New()
	// defer Pool.Close()
}

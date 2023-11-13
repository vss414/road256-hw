package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "gitlab.ozon.dev/vss414/hw-1/internal/counter"
	"gitlab.ozon.dev/vss414/hw-1/internal/validator"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

var SwaggerDir = "./swagger"

func main() {
	go rest()
	grpcServer()
}

func grpcServer() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, validator.New())

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func rest() {
	// Counter
	go func() {
		if err := http.ListenAndServe(":8089", nil); err != nil {
			log.Fatal(err)
		}
	}()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, gwmux, ":8082", opts); err != nil {
		log.Fatal(err)
	}

	// Serve the swagger-ui and swagger file
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	handleSwaggerFile(mux)

	// Register Swagger Handler
	fs := http.FileServer(http.Dir(SwaggerDir))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	server := &http.Server{Addr: ":8080", Handler: mux}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func handleSwaggerFile(mux *http.ServeMux) {
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		errorMessage := []byte("Failed to open swagger file.")
		f, err := os.Open("swagger/api/api.swagger.json")
		if err != nil {
			w.Write(errorMessage)
			return
		}

		_, err = io.Copy(w, f)
		if err != nil {
			w.Write(errorMessage)
			return
		}
	})
}

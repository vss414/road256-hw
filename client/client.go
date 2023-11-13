package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewAdminClient(clientConn)

	ctx := context.Background()

	in := bufio.NewReader(os.Stdin)

	for {
		var command string
		fmt.Fscan(in, &command)

		switch command {
		case "list":
			list(ctx, in, client)
		case "stream":
			stream(ctx, in, client)
		case "get":
			get(ctx, in, client)
		case "create":
			create(ctx, in, client)
		case "update":
			update(ctx, in, client)
		case "delete":
			deletePlayer(ctx, in, client)
		case "async-create":
			asyncCreate(ctx, in, client)
		case "async-update":
			asyncUpdate(ctx, in, client)
		case "async-delete":
			asyncDelete(ctx, in, client)
		case "pubsub-get":
			pubsubGet(ctx, in, client)
		case "pubsub-list":
			pubsubList(ctx, in, client)
		case "exit":
			return
		}
	}
}

func list(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var limit, page uint64
	var order, direction string
	fmt.Fscan(in, &limit, &page, &order, &direction)

	d, ok := pb.Direction_value[direction]
	if !ok {
		log.Println("wrong direction")
		return
	}

	o, ok := pb.Order_value[order]
	if !ok {
		log.Println("wrong order")
		return
	}

	response, err := client.PlayerList(ctx, &pb.PlayerListRequest{
		Limit:     limit,
		Page:      page,
		Order:     pb.Order(o),
		Direction: pb.Direction(d),
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("response: [%v]", response)
}

func pubsubList(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var limit, page uint64
	var order, direction string
	fmt.Fscan(in, &limit, &page, &order, &direction)

	d, ok := pb.Direction_value[direction]
	if !ok {
		log.Println("wrong direction")
		return
	}

	o, ok := pb.Order_value[order]
	if !ok {
		log.Println("wrong order")
		return
	}

	response, err := client.PlayerPubsubList(ctx, &pb.PlayerListRequest{
		Limit:     limit,
		Page:      page,
		Order:     pb.Order(o),
		Direction: pb.Direction(d),
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("response: [%v]", response)
}

func stream(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var limit, page uint64
	var order, direction string
	fmt.Fscan(in, &limit, &page, &order, &direction)

	d, ok := pb.Direction_value[direction]
	if !ok {
		log.Println("wrong direction")
		return
	}

	o, ok := pb.Order_value[order]
	if !ok {
		log.Println("wrong order")
		return
	}

	s, err := client.PlayerStreamList(ctx, &pb.PlayerStreamListRequest{
		Limit:     limit,
		Page:      page,
		Order:     pb.Order(o),
		Direction: pb.Direction(d),
	})
	if err != nil {
		log.Println(err)
		return
	}

	done := make(chan bool)

	go func() {
		for {
			r, err := s.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Printf("cannot receive %v", err)
			}
			log.Printf("Received player: %v", r)
		}
	}()

	<-done
	log.Printf("finished")
}

func get(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var id uint64
	fmt.Fscan(in, &id)

	response, err := client.PlayerGet(ctx, &pb.PlayerGetRequest{
		Id: id,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("response: [%v]", response)
}

func pubsubGet(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var id uint64
	fmt.Fscan(in, &id)

	response, err := client.PlayerPubsubGet(ctx, &pb.PlayerGetRequest{
		Id: id,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("response: [%v]", response)
}

func create(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var name, club string
	var games, goals, assists uint64
	fmt.Fscan(in, &name, &club, &games, &goals, &assists)

	_, err := client.PlayerCreate(ctx, &pb.PlayerCreateRequest{
		Name:    name,
		Club:    club,
		Games:   games,
		Goals:   goals,
		Assists: assists,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Done!")
}

func asyncCreate(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var name, club string
	var games, goals, assists uint64
	fmt.Fscan(in, &name, &club, &games, &goals, &assists)

	_, err := client.PlayerAsyncCreate(ctx, &pb.PlayerCreateRequest{
		Name:    name,
		Club:    club,
		Games:   games,
		Goals:   goals,
		Assists: assists,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Done!")
}

func update(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var name, club string
	var id, games, goals, assists uint64
	fmt.Fscan(in, &id, &name, &club, &games, &goals, &assists)

	_, err := client.PlayerUpdate(ctx, &pb.PlayerUpdateRequest{
		Id:      id,
		Name:    name,
		Club:    club,
		Games:   games,
		Goals:   goals,
		Assists: assists,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Done!")
}

func asyncUpdate(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var name, club string
	var id, games, goals, assists uint64
	fmt.Fscan(in, &id, &name, &club, &games, &goals, &assists)

	_, err := client.PlayerAsyncUpdate(ctx, &pb.PlayerUpdateRequest{
		Id:      id,
		Name:    name,
		Club:    club,
		Games:   games,
		Goals:   goals,
		Assists: assists,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Done!")
}

func deletePlayer(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var id uint64
	fmt.Fscan(in, &id)

	_, err := client.PlayerDelete(ctx, &pb.PlayerDeleteRequest{
		Id: id,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Done!")
}

func asyncDelete(ctx context.Context, in *bufio.Reader, client pb.AdminClient) {
	var id uint64
	fmt.Fscan(in, &id)

	_, err := client.PlayerAsyncDelete(ctx, &pb.PlayerDeleteRequest{
		Id: id,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Done!")
}

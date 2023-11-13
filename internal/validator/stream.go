package validator

import (
	"context"
	"gitlab.ozon.dev/vss414/hw-1/internal/counter"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"io"
	"log"
)

func (i *implementation) PlayerStreamList(in *pb.PlayerStreamListRequest, srv pb.Admin_PlayerStreamListServer) error {
	method := "stream list request"
	i.log.Infof("%s: %v", method, in)
	counter.PushInRequestsCounter()

	ctx := context.Background()

	s, err := i.client.PlayerStreamList(ctx, in)
	counter.PushOutRequestsCounter()

	if err != nil {
		i.log.Errorf("%s: client error: %s", method, err)
		counter.PushFailedRequestsCounter()
		return err
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
				i.log.Errorf("%s: receive error: %s", method, err)
				log.Fatalf("cannot receive %v", err)
			}
			if err := srv.Send(r); err != nil {
				i.log.Errorf("%s: send error: %s", method, err)
				log.Printf("send error %v", err)
			}
		}
	}()

	<-done

	i.log.Infof("%s: Done!", method)
	counter.PushSuccessRequestsCounter()

	return nil
}

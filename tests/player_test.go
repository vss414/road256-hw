//go:build integration
// +build integration

package tests

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"google.golang.org/protobuf/types/known/emptypb"
	"math"
	"testing"
	"time"
)

func TestPlayerCreate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		request, response := create(t)

		assert.IsType(t, &pb.PlayerCreateResponse{}, response)
		assert.Greater(t, response.Id, uint64(0))
		assert.Equal(t, request.Name, response.Name)
		assert.Equal(t, request.Club, response.Club)
		assert.Equal(t, request.Games, response.Games)
		assert.Equal(t, request.Goals, response.Goals)
		assert.Equal(t, request.Assists, response.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("wrong name", func(t *testing.T) {
			request := pb.PlayerCreateRequest{}

			response, err := Client.PlayerCreate(context.Background(), &request)

			expectedError := "rpc error: code = InvalidArgument desc = field: [name] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
			assert.Nil(t, response)
		})

		t.Run("wrong club", func(t *testing.T) {
			request := pb.PlayerCreateRequest{}
			err := faker.FakeData(&request.Name)
			require.NoError(t, err)

			response, err := Client.PlayerCreate(context.Background(), &request)

			expectedError := "rpc error: code = InvalidArgument desc = field: [club] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
			assert.Nil(t, response)
		})
	})
}

func TestPlayerAsyncCreate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		request := pb.PlayerCreateRequest{}
		err := faker.FakeData(&request)
		require.NoError(t, err)

		_, err = Client.PlayerAsyncCreate(context.Background(), &request)
		require.NoError(t, err)

		time.Sleep(time.Second)

		listRequest := pb.PlayerListRequest{
			Limit:     3,
			Page:      1,
			Order:     pb.Order_ORDER_ID,
			Direction: pb.Direction_DIRECTION_DESC,
		}

		response, err := Client.PlayerList(context.Background(), &listRequest)
		require.NoError(t, err)
		assert.Len(t, response.Players, int(listRequest.Limit))

		hasPlayer := false
		for _, p := range response.Players {
			if p.Name == request.Name {
				hasPlayer = true
			}
		}
		assert.True(t, hasPlayer)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("wrong name", func(t *testing.T) {
			request := pb.PlayerCreateRequest{}

			response, err := Client.PlayerAsyncCreate(context.Background(), &request)

			expectedError := "rpc error: code = InvalidArgument desc = field: [name] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
			assert.Nil(t, response)
		})

		t.Run("wrong club", func(t *testing.T) {
			request := pb.PlayerCreateRequest{}
			err := faker.FakeData(&request.Name)
			require.NoError(t, err)

			response, err := Client.PlayerAsyncCreate(context.Background(), &request)

			expectedError := "rpc error: code = InvalidArgument desc = field: [club] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
			assert.Nil(t, response)
		})
	})
}

func TestPlayerGet(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		_, response := create(t)

		player, err := Client.PlayerGet(context.Background(), &pb.PlayerGetRequest{
			Id: response.Id,
		})
		require.NoError(t, err)

		assert.Equal(t, response.Id, player.Id)
		assert.Equal(t, response.Name, player.Name)
		assert.Equal(t, response.Club, player.Club)
		assert.Equal(t, response.Games, player.Games)
		assert.Equal(t, response.Goals, player.Goals)
		assert.Equal(t, response.Assists, player.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		id := uint64(math.MaxInt32)
		_, err := Client.PlayerGet(context.Background(), &pb.PlayerGetRequest{
			Id: id,
		})
		expectedError := fmt.Sprintf(
			"rpc error: code = NotFound desc = player id: [%d]: player does not exist",
			id,
		)
		assert.Equal(t, expectedError, err.Error())
	})
}

func TestPlayerPubsubGet(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		_, response := create(t)

		player, err := Client.PlayerPubsubGet(context.Background(), &pb.PlayerGetRequest{
			Id: response.Id,
		})
		require.NoError(t, err)

		assert.Equal(t, response.Id, player.Id)
		assert.Equal(t, response.Name, player.Name)
		assert.Equal(t, response.Club, player.Club)
		assert.Equal(t, response.Games, player.Games)
		assert.Equal(t, response.Goals, player.Goals)
		assert.Equal(t, response.Assists, player.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		id := uint64(math.MaxInt32)
		_, err := Client.PlayerPubsubGet(context.Background(), &pb.PlayerGetRequest{
			Id: id,
		})
		expectedError := fmt.Sprintf(
			"rpc error: code = NotFound desc = player id: [%d]: player does not exist",
			id,
		)
		assert.Equal(t, expectedError, err.Error())
	})
}

func TestPlayerList(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("ASC direction", func(t *testing.T) {
			create(t)
			create(t)
			create(t)
			create(t)

			request := pb.PlayerListRequest{
				Limit:     3,
				Page:      1,
				Order:     pb.Order_ORDER_ID,
				Direction: pb.Direction_DIRECTION_ASC,
			}

			response, err := Client.PlayerList(context.Background(), &request)
			require.NoError(t, err)
			assert.Len(t, response.Players, int(request.Limit))

			for i, p := range response.Players {
				if i < len(response.Players)-1 {
					assert.Greater(t, response.Players[i+1].Id, p.Id)
				}
			}
		})

		t.Run("DESC direction", func(t *testing.T) {
			create(t)
			create(t)
			create(t)
			create(t)

			request := pb.PlayerListRequest{
				Limit:     3,
				Page:      1,
				Order:     pb.Order_ORDER_GAMES,
				Direction: pb.Direction_DIRECTION_DESC,
			}

			response, err := Client.PlayerList(context.Background(), &request)
			require.NoError(t, err)
			assert.Len(t, response.Players, int(request.Limit))

			for i, p := range response.Players {
				if i < len(response.Players)-1 {
					assert.GreaterOrEqual(t, p.Games, response.Players[i+1].Games)
				}
			}
		})
	})
}

func TestPlayerPubsubList(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("ASC direction", func(t *testing.T) {
			create(t)
			create(t)
			create(t)
			create(t)

			request := pb.PlayerListRequest{
				Limit:     3,
				Page:      1,
				Order:     pb.Order_ORDER_ID,
				Direction: pb.Direction_DIRECTION_ASC,
			}

			response, err := Client.PlayerPubsubList(context.Background(), &request)
			require.NoError(t, err)
			assert.Len(t, response.Players, int(request.Limit))

			for i, p := range response.Players {
				if i < len(response.Players)-1 {
					assert.Greater(t, response.Players[i+1].Id, p.Id)
				}
			}
		})

		t.Run("DESC direction", func(t *testing.T) {
			create(t)
			create(t)
			create(t)
			create(t)

			request := pb.PlayerListRequest{
				Limit:     3,
				Page:      1,
				Order:     pb.Order_ORDER_GAMES,
				Direction: pb.Direction_DIRECTION_DESC,
			}

			response, err := Client.PlayerPubsubList(context.Background(), &request)
			require.NoError(t, err)
			assert.Len(t, response.Players, int(request.Limit))

			for i, p := range response.Players {
				if i < len(response.Players)-1 {
					assert.GreaterOrEqual(t, p.Games, response.Players[i+1].Games)
				}
			}
		})
	})
}

func TestPlayerUpdate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		_, response := create(t)

		updateRequest := pb.PlayerUpdateRequest{}
		err := faker.FakeData(&updateRequest)
		require.NoError(t, err)
		updateRequest.Id = response.Id

		updateResponse, err := Client.PlayerUpdate(context.Background(), &updateRequest)
		require.NoError(t, err)
		assert.IsType(t, &emptypb.Empty{}, updateResponse)

		player, err := Client.PlayerGet(context.Background(), &pb.PlayerGetRequest{
			Id: response.Id,
		})
		require.NoError(t, err)

		assert.Equal(t, updateRequest.Id, player.Id)
		assert.Equal(t, updateRequest.Name, player.Name)
		assert.Equal(t, updateRequest.Club, player.Club)
		assert.Equal(t, updateRequest.Games, player.Games)
		assert.Equal(t, updateRequest.Goals, player.Goals)
		assert.Equal(t, updateRequest.Assists, player.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("wrong name", func(t *testing.T) {
			_, response := create(t)

			updateRequest := pb.PlayerUpdateRequest{}
			updateRequest.Id = response.Id

			updateResponse, err := Client.PlayerUpdate(context.Background(), &updateRequest)

			expectedError := "rpc error: code = InvalidArgument desc = field: [name] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
			assert.IsType(t, &emptypb.Empty{}, updateResponse)
		})

		t.Run("wrong club", func(t *testing.T) {
			_, response := create(t)

			updateRequest := pb.PlayerUpdateRequest{}
			updateRequest.Id = response.Id
			err := faker.FakeData(&updateRequest.Name)
			require.NoError(t, err)

			updateResponse, err := Client.PlayerUpdate(context.Background(), &updateRequest)

			expectedError := "rpc error: code = InvalidArgument desc = field: [club] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
			assert.IsType(t, &emptypb.Empty{}, updateResponse)
		})
	})
}

func TestPlayerAsyncUpdate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		_, response := create(t)

		updateRequest := pb.PlayerUpdateRequest{}
		err := faker.FakeData(&updateRequest)
		require.NoError(t, err)
		updateRequest.Id = response.Id

		updateResponse, err := Client.PlayerAsyncUpdate(context.Background(), &updateRequest)
		require.NoError(t, err)
		assert.IsType(t, &emptypb.Empty{}, updateResponse)

		time.Sleep(time.Second)

		player, err := Client.PlayerPubsubGet(context.Background(), &pb.PlayerGetRequest{
			Id: response.Id,
		})
		require.NoError(t, err)

		assert.Equal(t, updateRequest.Id, player.Id)
		assert.Equal(t, updateRequest.Name, player.Name)
		assert.Equal(t, updateRequest.Club, player.Club)
		assert.Equal(t, updateRequest.Games, player.Games)
		assert.Equal(t, updateRequest.Goals, player.Goals)
		assert.Equal(t, updateRequest.Assists, player.Assists)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("wrong name", func(t *testing.T) {
			_, response := create(t)

			updateRequest := pb.PlayerUpdateRequest{}
			updateRequest.Id = response.Id

			updateResponse, err := Client.PlayerAsyncUpdate(context.Background(), &updateRequest)

			expectedError := "rpc error: code = InvalidArgument desc = field: [name] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
			assert.IsType(t, &emptypb.Empty{}, updateResponse)
		})

		t.Run("wrong club", func(t *testing.T) {
			_, response := create(t)

			updateRequest := pb.PlayerUpdateRequest{}
			updateRequest.Id = response.Id
			err := faker.FakeData(&updateRequest.Name)
			require.NoError(t, err)

			updateResponse, err := Client.PlayerAsyncUpdate(context.Background(), &updateRequest)

			expectedError := "rpc error: code = InvalidArgument desc = field: [club] length should be between 0 and 30 symbols: invalid data"
			assert.Equal(t, expectedError, err.Error())
			assert.IsType(t, &emptypb.Empty{}, updateResponse)
		})
	})
}

func TestPlayerDelete(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		_, response := create(t)

		_, err := Client.PlayerDelete(context.Background(), &pb.PlayerDeleteRequest{
			Id: response.Id,
		})
		require.NoError(t, err)

		_, err = Client.PlayerGet(context.Background(), &pb.PlayerGetRequest{
			Id: response.Id,
		})
		expectedError := fmt.Sprintf(
			"rpc error: code = NotFound desc = player id: [%d]: player does not exist",
			response.Id,
		)
		assert.Equal(t, expectedError, err.Error())
	})

	t.Run("negative", func(t *testing.T) {
		id := uint64(math.MaxInt32)
		_, err := Client.PlayerDelete(context.Background(), &pb.PlayerDeleteRequest{
			Id: id,
		})
		expectedError := fmt.Sprintf(
			"rpc error: code = NotFound desc = player id: [%d]: player does not exist",
			id,
		)
		assert.Equal(t, expectedError, err.Error())
	})
}

func TestPlayerAsyncDelete(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		_, response := create(t)

		_, err := Client.PlayerAsyncDelete(context.Background(), &pb.PlayerDeleteRequest{
			Id: response.Id,
		})
		require.NoError(t, err)

		time.Sleep(time.Second)

		_, err = Client.PlayerPubsubGet(context.Background(), &pb.PlayerGetRequest{
			Id: response.Id,
		})
		expectedError := fmt.Sprintf(
			"rpc error: code = NotFound desc = player id: [%d]: player does not exist",
			response.Id,
		)
		assert.Equal(t, expectedError, err.Error())
	})
}

func create(t *testing.T) (*pb.PlayerCreateRequest, *pb.PlayerCreateResponse) {
	request := pb.PlayerCreateRequest{}
	err := faker.FakeData(&request)
	require.NoError(t, err)

	response, err := Client.PlayerCreate(context.Background(), &request)
	require.NoError(t, err)

	return &request, response
}

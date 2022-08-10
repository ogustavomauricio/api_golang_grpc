package main

import (
	"context"
	"testing"

	pb "gustavo/klever-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

func TestUpdate2(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCurrencyCoinServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Success to downvote", func(mt *mtest.T) {
		collection = mt.Coll
		fakeCoin := &CoinItem{
			Name: "BTC",
			Price: 1.300000,
			Vote: 2501,
		}

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
		})

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "name", Value: "BTC"},
			{Key: "price", Value: 1.300000},
			{Key: "vote", Value: 2505},
		}))

		req := &pb.CoinRequest{Name: fakeCoin.Name}
		_, err := c.DownvoteCoin(context.Background(), req)

		if err != nil {
			t.Errorf("Something went wrong: %v", err)
		}
	})
}

func TestNotExistedCoin2(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCurrencyCoinServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Success to downvote", func(mt *mtest.T) {
		collection = mt.Coll
		fakeCoin := &CoinItem{
			Name: "moedinha",
		}

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		req := &pb.CoinRequest{Name: fakeCoin.Name}
		_, err := c.DownvoteCoin(context.Background(), req)

		if err == nil {
			t.Errorf("Something went wrong: %v", err)
		}
	})
}

func TestRequiredNameUpdate2(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	client := pb.NewCurrencyCoinServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("It's not possible to downvote a coin when name is not passed", func(mt *mtest.T) {
		collection = mt.Coll

	req := &pb.CreateCoinRequest{}
	_, err := client.CreateCoin(context.Background(), req)

	if err == nil {
		t.Error("TestRequiredNameUpdate is not checking blank field name")
}
},
	)
}

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

func TestCreateCoin(t *testing.T) {
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

	mt.Run("It's possible to create a new coin", func(mt *mtest.T) {
		collection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		fakeCoin := &pb.CreateCoinRequest{
			Name:  "Ethereum",
			Price: 0.0006,
		}

		req := &pb.CreateCoinRequest{Name: fakeCoin.Name, Price: fakeCoin.Price}
		_, err := client.CreateCoin(context.Background(), req)

		if err != nil {
			t.Errorf("TestCreateCoin(%v) got unexpected error", err)
		}

	})
}

func TestCoinExisted(t *testing.T) {
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

	mt.Run("It's not possible to create a new coin when it already exists", func(mt *mtest.T) {
		collection = mt.Coll

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "name", Value: "Ethereum"},
			{Key: "price", Value: 1.300000},
			{Key: "vote", Value: 2501},
		}))

		fakeCoin := &pb.CreateCoinRequest{
			Name:  "Ethereum",
			Price: 0.0006,
		}

		req := &pb.CreateCoinRequest{Name: fakeCoin.Name, Price: fakeCoin.Price}
		_, err := client.CreateCoin(context.Background(), req)

		if err == nil {
			t.Error("CreateCoin is creating a repeated coin with existing name")
		}
	},
	)
}

func TestRequiredName(t *testing.T) {
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

	mt.Run("It's not possible to create a coin when name is not passed", func(mt *mtest.T) {
		collection = mt.Coll

		fakeCoin := &pb.CreateCoinRequest{
			Price: 0.0006,
		}

		req := &pb.CreateCoinRequest{Price: fakeCoin.Price}
		_, err := client.CreateCoin(context.Background(), req)

		if err == nil {
			t.Error("TestRequiredName is not checking blank field name")
		}
	},
	)
}

func TestRequiredPrice(t *testing.T) {
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

	mt.Run("It's not possible to create a coin when price is not passed", func(mt *mtest.T) {
		collection = mt.Coll

		fakeCoin := &pb.CreateCoinRequest{
			Name: "Ethereum",
		}

		req := &pb.CreateCoinRequest{Price: fakeCoin.Price}
		_, err := client.CreateCoin(context.Background(), req)

		if err == nil {
			t.Error("TestRequiredPrice is not checking blank field price")
		}
	},
	)
}

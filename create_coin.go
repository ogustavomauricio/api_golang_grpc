package main

import (
	"context"
	"fmt"

	pb "gustavo/klever-grpc/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkCoin(name string) error {
	result := &CoinItem{}
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	collection.FindOne(context.TODO(), filter).Decode(&result)

	if result.Name == name {
		return status.Error(
			codes.AlreadyExists,
			"Can't create coin with an existing name",
		)
	}

	return nil
}

func (s *Server) CreateCoin(ctx context.Context, in *pb.CreateCoinRequest) (*pb.CoinResponse, error) {

	data := &CoinItem{
		Name:  in.Name,
		Price: in.Price,
		Vote:  0,
	}
	err := data.CreateValidate()

	if err != nil {
		return nil, err
	}

	err = checkCoin(data.Name)

	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &pb.CoinResponse{
		Name:  data.Name,
		Price: data.Price,
		Vote:  data.Vote,
	}, nil
}

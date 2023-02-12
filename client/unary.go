package main

import (
	"context"
	"log"
	"time"

	pb "github.com/IRSHIT033/go-grpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not greet : %v", err)
	}
	log.Printf("%s", resp.Message)
}

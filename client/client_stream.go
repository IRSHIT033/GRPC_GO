package main

import (
	"context"
	"log"
	"time"

	pb "github.com/IRSHIT033/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("client streaming started ")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for _, name := range names.Message {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("send the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaing finished")
	if err != nil {
		log.Fatalf("error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}

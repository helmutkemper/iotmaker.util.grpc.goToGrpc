package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	var ci *pb.ContainerInspectReply
	var ni *pb.NetworkInspectReply

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDockerServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	ci, err = c.ContainerInspect(ctx, &pb.ContainerInspectRequest{ID: "0b85a0b07ad6ef643df3c6c4f91a728aa3c171cbcf067f961f13ed7e3702fb8b"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("ID: %s\n", ci.GetID())

	ni, err = c.NetworkInspect(ctx, &pb.NetworkInspectRequest{ID: "0fdb93c956338309125d1ff94e440448d1b70e2de29626260451bdda4b8d8226"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("ID: %s\n", ni.GetID())
}

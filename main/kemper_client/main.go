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
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDockerServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ContainerInspect(ctx, &pb.ContainerInspectRequest{ID: "666_server"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("ID: %s\nError: %+v\n", r.GetId(), r.GetError())
}

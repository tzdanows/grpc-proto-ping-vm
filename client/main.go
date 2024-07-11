package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-protobuf-ping-vm/ping/proto"

	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	vmIPAddress := os.Getenv("VM_IP_ADDR")
	if vmIPAddress == "" {
		log.Fatal("env variable not set")
	}

	conn, err := grpc.Dial(vmIPAddress+":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.PingRequest{Message: "Hello, Server!"})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}

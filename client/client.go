package main

import (
	"context"
	"log"
	"time"

	pb "github.com/khbdev/grpc-hello/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main(){
	    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("❌ Serverga ulana olmadim: %v", err)
    }
    defer conn.Close()


	client := pb.NewHelloServiceClient(conn)

	    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

	 resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Azizbek"})
    if err != nil {
        log.Fatalf("❌ RPC xato: %v", err)
    }
	    log.Printf("✅ Serverdan javob: %s", resp.Message)
		

}
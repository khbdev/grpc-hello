package main

import (
    "context"
    "log"
    "time"

    pb "github.com/khbdev/grpc-hello/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewHelloServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Azizbek"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Javob: %s", r.Message)
}

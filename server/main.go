package main

import (
    "context"
    "log"
    "net"

    pb "github.com/khbdev/grpc-hello/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{Message: "Salom, " + req.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("listen error: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterHelloServiceServer(s, &server{})

    log.Println("🚀 Server started on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("serve error: %v", err)
    }
}

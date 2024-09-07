package main

import (
	"context"
	pb "grpc-echo-server/echo"
	"log"
	"net"
	"net/http" // Import for HTTP server

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received message: %s", req.Message)
	return &pb.EchoResponse{Message: "Echo: " + req.Message}, nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register your service here
	// pb.RegisterEchoServiceServer(grpcServer, &server{})

	// Enable reflection
	reflection.Register(grpcServer)

	// Start gRPC server in a new goroutine
	go func() {
		log.Println("gRPC server is running on port :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Set up HTTP server for health checks
	http.HandleFunc("/healthz/ready", healthCheckHandler)
	log.Println("Health check server is running on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start health check server: %v", err)
	}
}

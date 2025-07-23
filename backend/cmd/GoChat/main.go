package GoChat

import (
	"GoChat/internal/infrastructure/db/postgres"
	"GoChat/internal/pb"
	"GoChat/internal/server"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	_ = godotenv.Load()

	db := postgres.Connect()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server.ChatServer{})
	log.Println("gRPC server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

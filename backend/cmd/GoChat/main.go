package main

import (
	"GoChat/internal/infrastructure/db/postgres"
	"GoChat/internal/pb"
	grpcserver "GoChat/internal/server/grpc"
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
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &grpcserver.ChatServer{
		DB: db,
	})

	log.Println("gRPC server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

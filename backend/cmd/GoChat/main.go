package main

import (
	"GoChat/internal/infrastructure/cache/redis"
	"GoChat/internal/infrastructure/db/postgres"
	"GoChat/internal/pb"
	grpcserver "GoChat/internal/server/grpc"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	db, err := postgres.Connect()
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	redis.InitRedis()
	if err != nil {
		log.Fatalln("Redis init error:", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcserver.AuthInterceptor(db)),
		grpc.StreamInterceptor(grpcserver.AuthStreamInterceptor(db)),
	)
	pb.RegisterChatServiceServer(s, &grpcserver.ChatServer{
		DB: db,
	})

	log.Println("gRPC server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

package main

import (
	"GoChat/internal/pb"
	"log"
	httplib "net/http"

	"GoChat/internal/infrastructure/http"
	"GoChat/internal/server/websocket"

	"google.golang.org/grpc"
)

func main() {
	httplib.HandleFunc("/ws", websocket.HandleWebSocket)

	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	grpcClient := pb.NewChatServiceClient(grpcConn)

	httplib.HandleFunc("/RegisterUser", http.RegisterUserHandler(grpcClient))

	log.Println("Server started at :8080")
	if err := httplib.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

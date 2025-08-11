package main

import (
	"GoChat/internal/pb"
	"log"
	httplib "net/http"

	"GoChat/internal/infrastructure/http"
	"GoChat/internal/server/websocket"

	"google.golang.org/grpc"
)

const gRPCPort = "localhost:50051"
const httpPort = ":8080"

func main() {
	httplib.HandleFunc("/ws", websocket.HandleWebSocket)

	grpcConn, err := grpc.Dial(gRPCPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	grpcClient := pb.NewChatServiceClient(grpcConn)

	httplib.HandleFunc("/RegisterUser", http.RegisterUserHandler(grpcClient))
	httplib.HandleFunc("/LoginUser", http.LoginUserHandler(grpcClient))

	log.Printf("Server started at %s", httpPort)
	if err := httplib.ListenAndServe(httpPort, nil); err != nil {
		log.Fatal(err)
	}
}

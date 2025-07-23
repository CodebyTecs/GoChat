package grpc

import (
	"GoChat/internal/pb"
	"GoChat/internal/server/websocket"
	"context"
	"log"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatServer) RegisterUser(ctx context.Context, user *pb.User) (*pb.Empty, error) {
	log.Printf("New user registered: %s\n", user.Username)
	return &pb.Empty{}, nil
}

func (s *ChatServer) SendMessage(ctx context.Context, msg *pb.Message) (*pb.Empty, error) {
	log.Printf("Message from %s to %s: %s", msg.Sender, msg.Receiver, msg.Text)

	websocket.MessageChannel <- msg
	return &pb.Empty{}, nil
}

func (s *ChatServer) StreamMessages(empty *pb.Empty, stream pb.ChatService_StreamMessagesServer) error {
	return nil
}

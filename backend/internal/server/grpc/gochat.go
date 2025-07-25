package grpc

import (
	"GoChat/internal/domain"
	"GoChat/internal/infrastructure/db/postgres"
	"GoChat/internal/pb"
	"GoChat/internal/server/websocket"
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	DB *sqlx.DB
}

func (s *ChatServer) RegisterUser(ctx context.Context, user *pb.User) (*pb.Empty, error) {
	domainUser := domain.User{
		Username: user.Username,
		Password: user.Password,
	}
	err := postgres.SaveUser(s.DB, domainUser)
	if err != nil {
		log.Printf("Failed to save user: %v", err)
		return nil, err
	}
	log.Printf("New user registered: %s\n", user.Username)
	return &pb.Empty{}, nil
}

func (s *ChatServer) SendMessage(ctx context.Context, msg *pb.Message) (*pb.Empty, error) {
	domainMessage := domain.Message{
		Sender:   msg.Sender,
		Receiver: msg.Receiver,
		Text:     msg.Text,
	}
	err := postgres.SaveMessage(s.DB, domainMessage)
	if err != nil {
		log.Printf("Failed to save message: %v", err)
		return nil, err
	}

	websocket.MessageChannel <- msg
	log.Printf("Message from %s to %s: %s", msg.Sender, msg.Receiver, msg.Text)
	return &pb.Empty{}, nil
}

func (s *ChatServer) StreamMessages(empty *pb.Empty, stream pb.ChatService_StreamMessagesServer) error {
	for {
		select {
		case msg := <-websocket.MessageChannel:
			if err := stream.Send(msg); err != nil {
				log.Printf("Failed to send message: %v", err)
				return err
			}
		case <-stream.Context().Done():
			log.Println("Client disconnected from stream")
			return nil
		}
	}
}

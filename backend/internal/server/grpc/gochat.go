package grpc

import (
	"GoChat/internal/domain"
	"GoChat/internal/infrastructure/cache/redis"
	"GoChat/internal/infrastructure/db/postgres"
	"GoChat/internal/pb"
	"GoChat/internal/server/websocket"
	"GoChat/pkg/auth"
	"context"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	DB *sqlx.DB
}

func (s *ChatServer) RegisterUser(ctx context.Context, user *pb.User) (*pb.TokenResponse, error) {
	domainUser := domain.User{
		Username: user.Username,
		Password: user.Password,
	}
	err := postgres.SaveUser(s.DB, domainUser)
	if err != nil {
		log.Printf("Failed to save user: %v", err)
		return nil, err
	}

	token, err := auth.GenerateToken(domainUser.Username)
	if err != nil {
		return nil, err
	}

	return &pb.TokenResponse{Token: token}, nil
}

func (s *ChatServer) LoginUser(ctx context.Context, user *pb.User) (*pb.TokenResponse, error) {
	dbUser, err := postgres.GetUserByUsername(s.DB, user.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	if dbUser.Password != user.Password {
		return nil, status.Error(codes.Unauthenticated, "Invalid password")
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "Token generation failed")
	}

	err = redis.Redis.Set(redis.Ctx, "jwt:"+token, user.Username, time.Hour).Err()
	if err != nil {
		log.Println("Redis error:", err)
	}

	return &pb.TokenResponse{Token: token}, nil
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

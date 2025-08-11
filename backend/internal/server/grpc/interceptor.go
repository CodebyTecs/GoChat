package grpc

import (
	"GoChat/internal/infrastructure/cache/redis"
	"GoChat/internal/infrastructure/db/postgres"
	"GoChat/pkg/auth"
	"context"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	redislib "github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(db *sqlx.DB) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("FullMethod:", info.FullMethod)
		if info.FullMethod == "/gochat.ChatService/RegisterUser" || info.FullMethod == "/gochat.ChatService/Login" {
			return handler(ctx, req)
		}

		username, err := getUsernameFromContext(ctx, db)
		if err != nil {
			log.Printf("Unauthorized access: %v", err)
			return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
		}

		log.Printf("Authorized username: %s", username)
		return handler(ctx, req)
	}
}

func AuthStreamInterceptor(db *sqlx.DB) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		_, err := getUsernameFromContext(ss.Context(), db)
		if err != nil {
			log.Printf("Stream unauthorized access: %v", err)
			return status.Error(codes.Unauthenticated, "Unauthorized")
		}
		return handler(srv, ss)
	}
}

func getUsernameFromContext(ctx context.Context, db *sqlx.DB) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("metadata not found")
	}

	tokens := md["authorization"]
	if len(tokens) == 0 {
		return "", errors.New("token missing")
	}

	token := tokens[0]

	username, err := redis.Redis.Get(redis.Ctx, "jwt:"+token).Result()
	if err == redislib.Nil {
		username, err = auth.ParseToken(token)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	exists, err := postgres.IsUserExist(db, username)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.New("user not found in DB")
	}

	return username, nil
}

package grpc

import (
	"GoChat/internal/infrastructure/cache/redis"
	"GoChat/pkg/auth"
	"context"
	"errors"
	redislib "github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)

func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		if strings.HasPrefix(info.FullMethod, "Login") || strings.HasPrefix(info.FullMethod, "RegisterUser") {
			return handler(ctx, req)
		}

		username, err := getUsernameFromContext(ctx)
		if err != nil {
			log.Printf("Unauthorized access: %v", err)
			return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
		}

		log.Printf("Authorized username: %s", username)
		return handler(ctx, req)
	}
}

func getUsernameFromContext(ctx context.Context) (string, error) {
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
		return auth.ParseToken(token)
	} else if err != nil {
		return "", err
	}
	return username, nil
}

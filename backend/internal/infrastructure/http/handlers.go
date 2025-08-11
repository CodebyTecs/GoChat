package http

import (
	"GoChat/internal/domain"
	"GoChat/internal/pb"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	}
}

func RegisterUserHandler(grpcClient pb.ChatServiceClient) http.HandlerFunc {
	return withCORS(func(w http.ResponseWriter, r *http.Request) {
		var req domain.UserRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		grpcResp, err := grpcClient.RegisterUser(context.Background(), &pb.User{
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			log.Println("RegisterUser gRPC error:", err)
			http.Error(w, "Registration failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": grpcResp.Token,
		})
	})
}

func LoginUserHandler(grpcClient pb.ChatServiceClient) http.HandlerFunc {
	return withCORS(func(w http.ResponseWriter, r *http.Request) {
		var req domain.UserRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		grpcResp, err := grpcClient.LoginUser(context.Background(), &pb.User{
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			log.Println("LoginUser gRPC error:", err)

			st, ok := status.FromError(err)
			if ok {
				switch st.Code() {
				case codes.NotFound:
					http.Error(w, "User not found", http.StatusNotFound)
					return
				case codes.Unauthenticated:
					http.Error(w, "Invalid password", http.StatusUnauthorized)
					return
				default:
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
				}
			}

			http.Error(w, "Unknown error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": grpcResp.Token,
		})
	})
}

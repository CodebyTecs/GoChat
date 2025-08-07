package http

import (
	"GoChat/internal/pb"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
		var req RegisterUserRequest

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

// server/server.go
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	pb "github.com/Khiem17204/go_recommendation_system/services/server/proto"
)

var (
	postgresConn = "postgres://user:password@postgres:5432/yourdb?sslmode=disable"
	proxyAddr    = "proxy:50051"
)

type RecRequest struct {
	Cards []string `json:"cards"`
}

type RecCardResponse struct {
	Cards []string `json:"cards"`
}

type RecDeckResponse struct {
	Decks []string `json:"decks"`
}

func main() {
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/rec_card", handleRecCard)
	http.HandleFunc("/rec_deck", handleRecDeck)

	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "query param 'q' is required", http.StatusBadRequest)
		return
	}
	db, err := sql.Open("postgres", postgresConn)
	if err != nil {
		http.Error(w, "database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, description FROM cards WHERE name ILIKE '%' || $1 || '%'", query)
	if err != nil {
		http.Error(w, "query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]string
	for rows.Next() {
		var name, desc string
		rows.Scan(&name, &desc)
		results = append(results, map[string]string{"name": name, "description": desc})
	}
	json.NewEncoder(w).Encode(results)
}

func handleRecCard(w http.ResponseWriter, r *http.Request) {
	var req RecRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(proxyAddr, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "grpc connection error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	client := pb.NewRecommendationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.CardRecommend(ctx, &pb.RecommendRequest{Cards: req.Cards})
	if err != nil {
		http.Error(w, "grpc call failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(RecCardResponse{Cards: res.Cards})
}

func handleRecDeck(w http.ResponseWriter, r *http.Request) {
	var req RecRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(proxyAddr, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "grpc connection error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	client := pb.NewRecommendationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.DeckRecommend(ctx, &pb.RecommendRequest{Cards: req.Cards})
	if err != nil {
		http.Error(w, "grpc call failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(RecDeckResponse{Decks: res.Decks})
}

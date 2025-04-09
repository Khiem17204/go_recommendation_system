// recommendation_system/main.go
package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/Khiem17204/go_recommendation_system/services/recommendation_system/proto"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

const (
	kafkaBroker       = "kafka:9092"
	cardRequestTopic  = "card_suggest_req"
	cardResponseTopic = "card_suggest_res"
	deckRequestTopic  = "deck_suggest_req"
	deckResponseTopic = "deck_suggest_res"
	cardVectorAddr    = "vectorstore-card:60051"
	deckVectorAddr    = "vectorstore-deck:60051"
)

type CardSuggestRequest struct {
	RequestID string   `json:"request_id"`
	Cards     []string `json:"cards"`
}

type CardSuggestResponse struct {
	RequestID string   `json:"request_id"`
	Cards     []string `json:"cards"`
}

type DeckSuggestResponse struct {
	RequestID string   `json:"request_id"`
	Decks     []string `json:"decks"`
}

func main() {
	go listenCardSuggest()
	go listenDeckSuggest()
	select {} // block forever
}

func listenCardSuggest() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   cardRequestTopic,
		GroupID: "recommendation-card",
	})
	w := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    cardResponseTopic,
		Balancer: &kafka.LeastBytes{},
	}
	conn, err := grpc.Dial(cardVectorAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to card vector store: %v", err)
	}
	defer conn.Close()
	client := pb.NewVectorServiceClient(conn)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Read error:", err)
			continue
		}
		var req CardSuggestRequest
		_ = json.Unmarshal(m.Value, &req)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		res, err := client.SearchSimilarCards(ctx, &pb.VectorRequest{Cards: req.Cards})
		if err != nil {
			log.Println("gRPC error:", err)
			continue
		}

		out := CardSuggestResponse{
			RequestID: req.RequestID,
			Cards:     res.Cards,
		}
		msg, _ := json.Marshal(out)
		_ = w.WriteMessages(context.Background(), kafka.Message{Value: msg})
	}
}

func listenDeckSuggest() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   deckRequestTopic,
		GroupID: "recommendation-deck",
	})
	w := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    deckResponseTopic,
		Balancer: &kafka.LeastBytes{},
	}
	conn, err := grpc.Dial(deckVectorAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to deck vector store: %v", err)
	}
	defer conn.Close()
	client := pb.NewVectorServiceClient(conn)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Read error:", err)
			continue
		}
		var req CardSuggestRequest
		_ = json.Unmarshal(m.Value, &req)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		res, err := client.SearchSimilarDecks(ctx, &pb.VectorRequest{Cards: req.Cards})
		if err != nil {
			log.Println("gRPC error:", err)
			continue
		}

		out := DeckSuggestResponse{
			RequestID: req.RequestID,
			Decks:     res.Decks,
		}
		msg, _ := json.Marshal(out)
		_ = w.WriteMessages(context.Background(), kafka.Message{Value: msg})
	}
}

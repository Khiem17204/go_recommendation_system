// proxy/main.go (Kafka Request/Response gRPC proxy)
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"

	pb "github.com/Khiem17204/go_recommendation_system/services/proxy/proto"
)

const (
	kafkaBroker     = "kafka:9092"
	cardReqTopic    = "card_suggest_req"
	cardResTopic    = "card_suggest_res"
	deckReqTopic    = "deck_suggest_req"
	deckResTopic    = "deck_suggest_res"
	responseTimeout = 5 * time.Second
)

type SuggestionServer struct {
	producer *kafka.Writer
	pb.UnimplementedRecommendationServiceServer
}

func (s *SuggestionServer) CardRecommend(ctx context.Context, req *pb.RecommendRequest) (*pb.CardResponse, error) {
	return s.sendKafkaAndWaitCard(ctx, cardReqTopic, cardResTopic, req)
}

func (s *SuggestionServer) DeckRecommend(ctx context.Context, req *pb.RecommendRequest) (*pb.DeckResponse, error) {
	return s.sendKafkaAndWaitDeck(ctx, deckReqTopic, deckResTopic, req)
}

func (s *SuggestionServer) sendKafkaAndWaitCard(ctx context.Context, reqTopic, resTopic string, req *pb.RecommendRequest) (*pb.CardResponse, error) {
	reqID := uuid.New().String()

	payload := map[string]interface{}{
		"request_id": reqID,
		"cards":      req.Cards,
	}
	data, _ := json.Marshal(payload)

	if err := s.producer.WriteMessages(ctx, kafka.Message{
		Topic: reqTopic,
		Value: data,
	}); err != nil {
		return nil, err
	}

	// Wait for Kafka response
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   resTopic,
		GroupID: "proxy-card-response-handler",
	})
	defer reader.Close()

	resCtx, cancel := context.WithTimeout(ctx, responseTimeout)
	defer cancel()

	for {
		m, err := reader.ReadMessage(resCtx)
		if err != nil {
			return nil, fmt.Errorf("timeout or error reading from Kafka: %w", err)
		}
		var res struct {
			RequestID string   `json:"request_id"`
			Cards     []string `json:"cards"`
		}
		_ = json.Unmarshal(m.Value, &res)
		if res.RequestID == reqID {
			return &pb.CardResponse{
				Status: "OK",
				Cards:  res.Cards,
			}, nil
		}
	}
}

func (s *SuggestionServer) sendKafkaAndWaitDeck(ctx context.Context, reqTopic, resTopic string, req *pb.RecommendRequest) (*pb.DeckResponse, error) {
	reqID := uuid.New().String()

	payload := map[string]interface{}{
		"request_id": reqID,
		"cards":      req.Cards,
	}
	data, _ := json.Marshal(payload)

	if err := s.producer.WriteMessages(ctx, kafka.Message{
		Topic: reqTopic,
		Value: data,
	}); err != nil {
		return nil, err
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   resTopic,
		GroupID: "proxy-deck-response-handler",
	})
	defer reader.Close()

	resCtx, cancel := context.WithTimeout(ctx, responseTimeout)
	defer cancel()

	for {
		m, err := reader.ReadMessage(resCtx)
		if err != nil {
			return nil, fmt.Errorf("timeout or error reading from Kafka: %w", err)
		}
		var res struct {
			RequestID string   `json:"request_id"`
			Decks     []string `json:"decks"`
		}
		_ = json.Unmarshal(m.Value, &res)
		if res.RequestID == reqID {
			return &pb.DeckResponse{
				Status: "OK",
				Decks:  res.Decks,
			}, nil
		}
	}
}

func main() {
	producer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Balancer: &kafka.LeastBytes{},
	}
	defer producer.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(grpcServer, &SuggestionServer{producer: producer})

	log.Println("gRPC proxy server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"odinfleet-nakama/api"
)

const (
	defaultListenAddr = ":8080"
	defaultGamePort   = 7777
)

type server struct {
	client    *api.APIClient
	serviceID int32
	gamePort  int32
}

type startMatchRequest struct {
	MatchID  string   `json:"match_id"`
	GameMode string   `json:"game_mode"`
	MapGUID  string   `json:"map_guid"`
	JoinMode string   `json:"join_mode"`
	UserIDs  []string `json:"user_ids"`
}

type finishMatchRequest struct {
	MatchID string `json:"match_id"`
	Outcome string `json:"outcome"`
}

func main() {
	cfg, serviceID, gamePort, listenAddr, err := loadConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	srv := &server{
		client:    api.NewAPIClient(cfg),
		serviceID: serviceID,
		gamePort:  gamePort,
	}

	if err := srv.setReady(context.Background()); err != nil {
		log.Fatalf("set ready: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", srv.handleHealth)
	mux.HandleFunc("/match/start", srv.handleStartMatch)
	mux.HandleFunc("/match/finish", srv.handleFinishMatch)

	log.Printf("mock game server listening on %s", listenAddr)
	if err := http.ListenAndServe(listenAddr, mux); err != nil {
		log.Fatal(err)
	}
}

func loadConfigFromEnv() (*api.Configuration, int32, int32, string, error) {
	apiToken := os.Getenv("ODIN_API_TOKEN")
	if apiToken == "" {
		return nil, 0, 0, "", errors.New("missing ODIN_API_TOKEN")
	}

	serviceIDRaw := os.Getenv("ODIN_SERVICE_ID")
	if serviceIDRaw == "" {
		return nil, 0, 0, "", errors.New("missing ODIN_SERVICE_ID")
	}
	serviceID, err := strconv.Atoi(serviceIDRaw)
	if err != nil {
		return nil, 0, 0, "", fmt.Errorf("invalid ODIN_SERVICE_ID: %w", err)
	}

	gamePort := defaultGamePort
	if portRaw := os.Getenv("ODIN_GAME_PORT"); portRaw != "" {
		parsedPort, err := strconv.Atoi(portRaw)
		if err != nil {
			return nil, 0, 0, "", fmt.Errorf("invalid ODIN_GAME_PORT: %w", err)
		}
		gamePort = parsedPort
	}

	listenAddr := os.Getenv("MOCK_SERVER_ADDR")
	if listenAddr == "" {
		listenAddr = defaultListenAddr
	}

	cfg := api.NewConfiguration()
	cfg.Host = "fleet.4players.io"
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %v", apiToken))

	return cfg, int32(serviceID), int32(gamePort), listenAddr, nil
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

func (s *server) handleStartMatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not allowed"})
		return
	}

	var req startMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid JSON body"})
		return
	}
	if req.MatchID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "match_id is required"})
		return
	}

	if req.JoinMode == "" {
		req.JoinMode = "free"
	}

	metadata := map[string]any{
		"instance_state": "occupied",
		"match_id":       req.MatchID,
		"game_mode":      req.GameMode,
		"map_guid":       req.MapGUID,
		"join_mode":      req.JoinMode,
		"user_ids":       req.UserIDs,
		"game_port":      s.gamePort,
	}
	if err := s.updateMetadata(r.Context(), metadata); err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{"status": "match started"})
}

func (s *server) handleFinishMatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not allowed"})
		return
	}

	var req finishMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid JSON body"})
		return
	}

	metadata := map[string]any{
		"instance_state": "idle",
		"match_id":       "",
		"last_outcome":   req.Outcome,
		"user_ids":       []string{},
		"game_port":      s.gamePort,
	}
	if err := s.updateMetadata(r.Context(), metadata); err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{"status": "match finished"})
}

func (s *server) setReady(ctx context.Context) error {
	metadata := map[string]any{
		"instance_state": "idle",
		"game_port":      s.gamePort,
	}
	return s.updateMetadata(ctx, metadata)
}

func (s *server) updateMetadata(ctx context.Context, metadata map[string]any) error {
	req := s.client.DockerAPI.DockerServicesMetadataUpdate(ctx, s.serviceID).
		PatchMetadataRequest(api.PatchMetadataRequest{Metadata: metadata})
	_, _, err := req.Execute()
	return err
}

func writeJSON(w http.ResponseWriter, status int, payload map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

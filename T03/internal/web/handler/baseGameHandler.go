package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"tictactoe/internal/domain/model"
	webservice "tictactoe/internal/web/service"

	"github.com/google/uuid"
)

type BaseGameHandler struct {
	service webservice.WebService
}

func NewBaseGameHandler(service webservice.WebService) GameHandler {
	return &BaseGameHandler{service: service}
}

func (h *BaseGameHandler) CreateGameHumanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	game := h.service.NewGameHuman()
	gameInfo, _ := json.Marshal(game)
	w.Write([]byte(gameInfo))
}

func (h *BaseGameHandler) CreateGameAiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	game := h.service.NewGameAi()
	gameInfo, _ := json.Marshal(game)
	w.Write([]byte(gameInfo))
}

func (h *BaseGameHandler) NextMoveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlSplit := strings.Split(r.URL.Path, "/")
	if len(urlSplit) < 3 {
		http.Error(w, "game ID not found", http.StatusBadRequest)
		return
	}
	gameIdStr := urlSplit[2]
	gameId, err := uuid.Parse(gameIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	var coordinate model.Coordinate
	err = json.Unmarshal(body, &coordinate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game, err := h.service.NextMove(gameId, coordinate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gameInfo, _ := json.Marshal(game)
	w.Write([]byte(gameInfo))

}

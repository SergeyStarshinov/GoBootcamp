package handler

import "net/http"

type GameHandler interface {
	NextMoveHandler(w http.ResponseWriter, r *http.Request)
	CreateGameHumanHandler(w http.ResponseWriter, r *http.Request)
	CreateGameAiHandler(w http.ResponseWriter, r *http.Request)
}

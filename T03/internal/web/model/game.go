package webmodel

import "github.com/google/uuid"

type Game struct {
	GameId    uuid.UUID
	GameBoard GameBoard
	Message   string
}

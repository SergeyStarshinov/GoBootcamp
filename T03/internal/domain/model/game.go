package model

import "github.com/google/uuid"

type Game struct {
	GameBoard GameBoard
	GameId    uuid.UUID
}

func NewGame() Game {
	return Game{
		GameBoard: NewGameBoard(),
		GameId:    uuid.New(),
	}
}

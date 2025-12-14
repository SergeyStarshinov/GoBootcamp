package gameservice

import (
	"tictactoe/internal/domain/model"

	"github.com/google/uuid"
)

type GameService interface {
	NextAiMove(gameBoard model.GameBoard) model.GameBoard
	IsValid(gameBoard model.GameBoard, coordinate model.Coordinate) error
	IsFinished(model.GameBoard) int
	NewGame(startAI bool) model.Game
	SaveGame(game model.Game)
	LoadGame(gameId uuid.UUID) (model.Game, error)
}

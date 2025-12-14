package webservice

import (
	"tictactoe/internal/domain/model"
	webmodel "tictactoe/internal/web/model"

	"github.com/google/uuid"
)

type WebService interface {
	NewGameAi() webmodel.Game
	NewGameHuman() webmodel.Game
	NextMove(gameId uuid.UUID, coordinate model.Coordinate) (webmodel.Game, error)
}

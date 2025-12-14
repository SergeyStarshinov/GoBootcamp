package dataservice

import (
	"github.com/google/uuid"

	"tictactoe/internal/domain/model"
)

type DataService interface {
	SaveGame(model.Game)
	LoadGame(uuid uuid.UUID) (model.Game, error)
}

package repository

import (
	"github.com/google/uuid"

	datamodel "tictactoe/internal/datasource/model"
)

type DataRepository interface {
	SaveGame(uuid uuid.UUID, gameBoard datamodel.GameBoardDTO)
	LoadGame(uuid uuid.UUID) (datamodel.GameBoardDTO, error)
}

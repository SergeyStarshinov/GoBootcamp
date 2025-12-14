package repository

import (
	"fmt"

	"github.com/google/uuid"

	datamodel "tictactoe/internal/datasource/model"
)

type InMemoryRepository struct {
	store *datamodel.Storage
}

func NewInMemoryRepository(store *datamodel.Storage) DataRepository {
	return &InMemoryRepository{store: store}
}

func (r *InMemoryRepository) SaveGame(uuid uuid.UUID, gameBoard datamodel.GameBoardDTO) {
	r.store.Data.Store(uuid, gameBoard)
}

func (r *InMemoryRepository) LoadGame(uuid uuid.UUID) (datamodel.GameBoardDTO, error) {
	if value, ok := r.store.Data.Load(uuid); ok {
		return value.(datamodel.GameBoardDTO), nil
	}
	return datamodel.GameBoardDTO{}, fmt.Errorf("game with uuid %s not found", uuid.String())
}

package dataservice

import (
	"fmt"

	datamappers "tictactoe/internal/datasource/mappers"
	"tictactoe/internal/datasource/repository"
	"tictactoe/internal/domain/model"

	"github.com/google/uuid"
)

type BaseDataService struct {
	repo repository.DataRepository
}

func NewBaseDataService(r repository.DataRepository) DataService {
	return &BaseDataService{repo: r}
}

func (s BaseDataService) SaveGame(game model.Game) {
	s.repo.SaveGame(game.GameId, datamappers.FromGameBoardToDTO(game.GameBoard))
}

func (s BaseDataService) LoadGame(uuid uuid.UUID) (model.Game, error) {
	if gameBoardDTO, err := s.repo.LoadGame(uuid); err == nil {
		return model.Game{GameBoard: datamappers.FromDTOToGameBoard(gameBoardDTO), GameId: uuid}, nil
	} else {
		return model.NewGame(), fmt.Errorf("loading error: %w", err)
	}
}

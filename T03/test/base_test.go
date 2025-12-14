package test

import (
	"testing"

	datamappers "tictactoe/internal/datasource/mappers"
	datamodel "tictactoe/internal/datasource/model"
	"tictactoe/internal/datasource/repository"
	dataservice "tictactoe/internal/datasource/service"
	"tictactoe/internal/domain/model"
	"tictactoe/internal/infrastructure/constants"

	"github.com/google/uuid"
)

func TestCreateGameAiStart(t *testing.T) {
	game := minMaxService.NewGame(true)
	expectedAI := 1
	expectedVoid := 8

	marksCount, voidCount := countMarks(game)
	if expectedAI != marksCount {
		t.Errorf("Result was incorrect, got: %d AI marks, want: %d.", marksCount, expectedAI)
	}
	if expectedVoid != voidCount {
		t.Errorf("Result was incorrect, got: %d void field, want: %d.", voidCount, expectedVoid)
	}
}

func TestCreateGamePlayerStart(t *testing.T) {
	game := minMaxService.NewGame(false)
	expectedAI := 0
	expectedVoid := 9

	marksCount, voidCount := countMarks(game)
	if expectedAI != marksCount {
		t.Errorf("Result was incorrect, got: %d AI marks, want: %d.", marksCount, expectedAI)
	}
	if expectedVoid != voidCount {
		t.Errorf("Result was incorrect, got: %d void field, want: %d.", voidCount, expectedVoid)
	}
}

func TestRepositorySaveLoadGame(t *testing.T) {
	game := minMaxService.NewGame(true)
	testUuid := game.GameId
	repo := repository.NewInMemoryRepository(datamodel.NewStorage())
	repo.SaveGame(game.GameId, datamappers.FromGameBoardToDTO(game.GameBoard))
	if loadedGameBoard, err := repo.LoadGame(testUuid); err == nil {
		loadedGame := model.Game{GameBoard: datamappers.FromDTOToGameBoard(loadedGameBoard), GameId: testUuid}
		if loadedGame != game {
			t.Errorf("loaded game is not saved game")
		}
	} else {
		t.Errorf("load error: %s\n", err)
	}
}

func TestServiceSaveLoadGame(t *testing.T) {
	game := minMaxService.NewGame(false)
	testUuid := game.GameId
	repo := repository.NewInMemoryRepository(datamodel.NewStorage())
	service := dataservice.NewBaseDataService(repo)
	service.SaveGame(game)
	if loadedGame, err := service.LoadGame(testUuid); err == nil {
		if loadedGame != game {
			t.Errorf("loaded game is not saved game")
		}
	} else {
		t.Errorf("loading error: %s\n", err)
	}
	if _, err := service.LoadGame(uuid.New()); err == nil {
		t.Errorf("loading a non-existent game")
	}
}

func countMarks(game model.Game) (int, int) {
	marksCount := 0
	voidCount := 0
	for row := range constants.FIELD_SIZE {
		for col := range constants.FIELD_SIZE {
			mark, _ := game.GameBoard.Cell(row, col)
			switch mark {
			case constants.AI_PLAYER:
				marksCount++
			case constants.VOID:
				voidCount++
			}
		}
	}
	return marksCount, voidCount
}

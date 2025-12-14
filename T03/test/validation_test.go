package test

import (
	"testing"

	"tictactoe/internal/domain/model"
	"tictactoe/internal/infrastructure/constants"

	"github.com/google/uuid"
)

func TestInvalidCoordinate(t *testing.T) {
	game := minMaxService.NewGame(false)
	testCases := []model.Coordinate{
		model.NewCoordinate(-1, 0),
		model.NewCoordinate(1, -2),
		model.NewCoordinate(-1, -1),
		model.NewCoordinate(constants.FIELD_SIZE+1, 0),
		model.NewCoordinate(0, constants.FIELD_SIZE+1),
		model.NewCoordinate(constants.FIELD_SIZE+1, constants.FIELD_SIZE+2),
		model.NewCoordinate(-1, constants.FIELD_SIZE+1),
	}
	for i, coordinate := range testCases {
		if err := minMaxService.IsValid(game.GameBoard, coordinate); err == nil {
			t.Errorf("test %d with invalid coordinate. got: nil, expected: error.", i)
		}
	}
}

func TestValidationGameOccupied(t *testing.T) {
	ai := constants.AI_PLAYER
	human := constants.HUMAN_PLAYER
	void := constants.VOID

	game := model.Game{
		GameBoard: model.GameBoard{
			{ai, human, ai},
			{human, ai, human},
			{human, ai, void},
		},
		GameId: uuid.New(),
	}
	ds.SaveGame(game)

	testCases := []model.Coordinate{
		model.NewCoordinate(0, 0),
		model.NewCoordinate(0, 1),
		model.NewCoordinate(1, 1),
		model.NewCoordinate(2, 0),
		model.NewCoordinate(1, 1),
	}

	for i, coordinate := range testCases {
		if err := minMaxService.IsValid(game.GameBoard, coordinate); err == nil {
			t.Errorf("test %d with occupied coordinate. got: nil, expected: error.", i)
		}
	}
}

func TestValidationGameNotOccupied(t *testing.T) {

	game := minMaxService.NewGame(false)
	ds.SaveGame(game)

	testCases := []model.Coordinate{
		model.NewCoordinate(0, 0),
		model.NewCoordinate(0, 1),
		model.NewCoordinate(1, 1),
		model.NewCoordinate(2, 0),
		model.NewCoordinate(1, 1),
	}

	for i, coordinate := range testCases {
		if err := minMaxService.IsValid(game.GameBoard, coordinate); err != nil {
			t.Errorf("test %d with occupied coordinate. got: error, expected: nil.", i)
		}
	}
}

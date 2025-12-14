package test

import (
	"testing"

	datamodel "tictactoe/internal/datasource/model"
	"tictactoe/internal/datasource/repository"
	dataservice "tictactoe/internal/datasource/service"
	"tictactoe/internal/domain/model"
	gameservice "tictactoe/internal/domain/service"
	"tictactoe/internal/infrastructure/constants"
)

var r = repository.NewInMemoryRepository(datamodel.NewStorage())
var ds = dataservice.NewBaseDataService(r)
var minMaxService = gameservice.NewMinMax(ds)

func TestEndGamePlayerWin(t *testing.T) {
	ai := constants.AI_PLAYER
	human := constants.HUMAN_PLAYER
	void := constants.VOID

	testCases := []model.GameBoard{
		{
			{human, human, human},
			{human, ai, human},
			{ai, human, ai},
		},
		{
			{human, ai, human},
			{human, human, human},
			{ai, human, ai},
		},
		{
			{ai, ai, human},
			{human, ai, ai},
			{human, human, human},
		},
		{
			{human, ai, human},
			{human, ai, ai},
			{human, void, void},
		},
		{
			{ai, human, human},
			{ai, human, ai},
			{void, human, void},
		},
		{
			{ai, human, human},
			{ai, ai, human},
			{void, void, human},
		},
		{
			{human, ai, human},
			{ai, human, ai},
			{void, void, human},
		},
		{
			{human, ai, human},
			{ai, human, ai},
			{void, void, human},
		},
	}

	for _, gameBoard := range testCases {
		if winner := minMaxService.IsFinished(gameBoard); winner != human {
			t.Errorf("winner is: %d, expected: %d.", winner, human)
		}
	}
}

func TestEndGameAiWin(t *testing.T) {
	ai := constants.AI_PLAYER
	human := constants.HUMAN_PLAYER
	void := constants.VOID

	testCases := []model.GameBoard{
		{
			{ai, ai, ai},
			{human, void, human},
			{ai, human, void},
		},
		{
			{human, ai, human},
			{ai, ai, ai},
			{human, human, void},
		},
		{
			{human, ai, human},
			{human, void, human},
			{ai, ai, ai},
		},
		{
			{ai, human, human},
			{ai, human, ai},
			{ai, void, void},
		},
		{
			{human, ai, human},
			{human, ai, ai},
			{void, ai, void},
		},
		{
			{ai, human, ai},
			{human, human, ai},
			{human, void, ai},
		},
		{
			{ai, human, human},
			{human, ai, human},
			{void, void, ai},
		},
		{
			{human, human, ai},
			{human, ai, ai},
			{ai, void, human},
		},
	}

	for _, gameBoard := range testCases {
		if winner := minMaxService.IsFinished(gameBoard); winner != ai {
			t.Errorf("winner is: %d, expected: %d.", winner, ai)
		}
	}
}

func TestEndGameNotEnd(t *testing.T) {
	ai := constants.AI_PLAYER
	human := constants.HUMAN_PLAYER
	void := constants.VOID

	testCases := []model.GameBoard{
		{
			{ai, human, ai},
			{human, ai, human},
			{human, ai, void},
		},
		{
			{void, void, void},
			{void, human, void},
			{ai, void, void},
		},
		{
			{void, void, void},
			{void, void, void},
			{void, void, void},
		},
	}

	for _, gameBoard := range testCases {
		if winner := minMaxService.IsFinished(gameBoard); winner != void {
			t.Errorf("winner is: %d, expected: %d.", winner, void)
		}
	}
}

func TestEndGameDraw(t *testing.T) {
	ai := constants.AI_PLAYER
	human := constants.HUMAN_PLAYER

	testCases := []model.GameBoard{
		{
			{ai, human, ai},
			{human, ai, human},
			{human, ai, human},
		},
		{
			{human, ai, human},
			{ai, human, human},
			{ai, human, ai},
		},
	}

	for _, gameBoard := range testCases {
		if winner := minMaxService.IsFinished(gameBoard); winner != constants.DRAW {
			t.Errorf("winner is: %d, expected: %d.", winner, constants.DRAW)
		}
	}
}

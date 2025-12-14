package webservice

import (
	"fmt"
	"tictactoe/internal/domain/model"
	gameservice "tictactoe/internal/domain/service"
	"tictactoe/internal/infrastructure/constants"
	webmappers "tictactoe/internal/web/mappers"
	webmodel "tictactoe/internal/web/model"

	"github.com/google/uuid"
)

type BaseWebService struct {
	gs gameservice.GameService
}

func NewBaseWebService(gs gameservice.GameService) WebService {
	return &BaseWebService{gs: gs}
}

func (ws *BaseWebService) NewGameAi() webmodel.Game {
	return webmappers.FromGameToWeb(ws.gs.NewGame(true))
}

func (ws *BaseWebService) NewGameHuman() webmodel.Game {
	return webmappers.FromGameToWeb(ws.gs.NewGame(false))
}

func (ws *BaseWebService) NextMove(gameId uuid.UUID, coordinate model.Coordinate) (webmodel.Game, error) {
	game, err := ws.gs.LoadGame(gameId)
	if err != nil {
		return ws.NewGameHuman(), err
	}

	if winner := ws.gs.IsFinished(game.GameBoard); winner != constants.VOID {
		return ws.NewGameHuman(), fmt.Errorf("this game is over")
	}

	if err := ws.gs.IsValid(game.GameBoard, coordinate); err != nil {
		return ws.NewGameHuman(), fmt.Errorf("invalid move: %w", err)
	}

	game.GameBoard.SetCell(coordinate, constants.HUMAN_PLAYER)
	if winner := ws.gs.IsFinished(game.GameBoard); winner != constants.VOID {
		ws.gs.SaveGame(game)
		gameWeb := webmappers.FromGameToWeb(game)
		if winner == constants.HUMAN_PLAYER {
			gameWeb.Message = "Congratulations, you won!!!"
		} else {
			gameWeb.Message = "This game ended in a draw"
		}
		return gameWeb, nil
	}

	game.GameBoard = ws.gs.NextAiMove(game.GameBoard)
	ws.gs.SaveGame(game)
	gameWeb := webmappers.FromGameToWeb(game)
	if winner := ws.gs.IsFinished(game.GameBoard); winner != constants.VOID {
		if winner == constants.AI_PLAYER {
			gameWeb.Message = "Game over. AI won!!!"
		} else {
			gameWeb.Message = "This game ended in a draw"
		}
	} else {
		gameWeb.Message = "Make your next move"
	}
	return gameWeb, nil
}

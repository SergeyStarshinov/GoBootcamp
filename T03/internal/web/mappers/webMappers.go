package webmappers

import (
	"tictactoe/internal/domain/model"
	"tictactoe/internal/infrastructure/constants"
	webmodel "tictactoe/internal/web/model"
)

func FromGameBoardToWeb(gb model.GameBoard) webmodel.GameBoard {
	var gameBoard webmodel.GameBoard
	for row := range constants.FIELD_SIZE {
		for col := range constants.FIELD_SIZE {
			switch gb[row][col] {
			case constants.AI_PLAYER:
				gameBoard[row][col] = constants.AI_SIGN
			case constants.HUMAN_PLAYER:
				gameBoard[row][col] = constants.HUMAN_SIGN
			default:
				gameBoard[row][col] = constants.VOID_SIGN
			}
		}
	}
	return gameBoard
}

func FromGameToWeb(g model.Game) webmodel.Game {
	return webmodel.Game{
		GameId:    g.GameId,
		GameBoard: FromGameBoardToWeb(g.GameBoard),
		Message:   "",
	}
}

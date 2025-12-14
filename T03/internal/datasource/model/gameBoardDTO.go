package datamodel

import "tictactoe/internal/infrastructure/constants"

type GameBoardDTO struct {
	Board [constants.FIELD_SIZE * constants.FIELD_SIZE]int
}

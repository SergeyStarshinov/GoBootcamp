package datamappers

import (
	datamodel "tictactoe/internal/datasource/model"
	"tictactoe/internal/domain/model"
	"tictactoe/internal/infrastructure/constants"
)

func FromGameBoardToDTO(gb model.GameBoard) datamodel.GameBoardDTO {
	var gbDTO datamodel.GameBoardDTO
	for row := range constants.FIELD_SIZE {
		for col := range constants.FIELD_SIZE {
			gbDTO.Board[row*constants.FIELD_SIZE+col] = gb[row][col]
		}
	}
	return gbDTO
}

func FromDTOToGameBoard(gbDTO datamodel.GameBoardDTO) model.GameBoard {
	var gb model.GameBoard
	for row := range constants.FIELD_SIZE {
		for col := range constants.FIELD_SIZE {
			gb[row][col] = gbDTO.Board[row*constants.FIELD_SIZE+col]
		}
	}
	return gb
}

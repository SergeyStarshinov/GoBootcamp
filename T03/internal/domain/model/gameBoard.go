package model

import (
	"fmt"

	"tictactoe/internal/infrastructure/constants"
)

type GameBoard [constants.FIELD_SIZE][constants.FIELD_SIZE]int

func NewGameBoard() GameBoard {
	var gameBoard GameBoard
	for row := range constants.FIELD_SIZE {
		for col := range constants.FIELD_SIZE {
			gameBoard[row][col] = constants.VOID
		}
	}
	return gameBoard
}

func (g GameBoard) CellByCoordinate(coordinate Coordinate) (int, error) {
	if err := coordinate.isValid(); err != nil {
		return constants.VOID, fmt.Errorf("invalid coordinate: %w", err)
	}
	return g[coordinate.Row][coordinate.Col], nil
}

func (g GameBoard) Cell(row, col int) (int, error) {
	coordinate := NewCoordinate(row, col)
	return g.CellByCoordinate(coordinate)
}

func (g *GameBoard) SetCell(coordinate Coordinate, value int) error {
	if err := coordinate.isValid(); err != nil {
		return fmt.Errorf("invalid coordinate: %w", err)
	}
	g[coordinate.Row][coordinate.Col] = value
	return nil
}

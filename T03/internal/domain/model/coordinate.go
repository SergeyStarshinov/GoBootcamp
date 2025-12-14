package model

import (
	"fmt"

	"tictactoe/internal/infrastructure/constants"
)

type Coordinate struct {
	Row int
	Col int
}

func NewCoordinate(row, col int) Coordinate {
	return Coordinate{Row: row, Col: col}
}

func (c Coordinate) isValid() error {
	if c.Row >= constants.FIELD_SIZE || c.Col >= constants.FIELD_SIZE {
		return fmt.Errorf("coordinate exceeds maximum")
	}
	if c.Row < 0 || c.Col < 0 {
		return fmt.Errorf("coordinate less than 0")
	}
	return nil
}

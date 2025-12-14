package gameservice

import (
	"fmt"
	"math/rand/v2"
	"slices"

	dataservice "tictactoe/internal/datasource/service"
	"tictactoe/internal/domain/model"
	"tictactoe/internal/infrastructure/constants"

	"github.com/google/uuid"
)

type Move struct {
	coordinate model.Coordinate
	score      int
}

type MinMaxService struct {
	dataService dataservice.DataService
}

func NewMinMax(dataService dataservice.DataService) GameService {
	return &MinMaxService{dataService: dataService}
}

func (m *MinMaxService) NextAiMove(gameBoard model.GameBoard) model.GameBoard {
	nextMove := m.minMax(gameBoard, constants.AI_PLAYER).coordinate
	gameBoard.SetCell(nextMove, constants.AI_PLAYER)
	return gameBoard
}

func (m *MinMaxService) NewGame(startAI bool) model.Game {
	game := model.NewGame()
	if startAI {
		row := rand.IntN(constants.FIELD_SIZE)
		col := rand.IntN(constants.FIELD_SIZE)
		game.GameBoard.SetCell(model.NewCoordinate(row, col), constants.AI_PLAYER)
	}
	m.dataService.SaveGame(game)
	return game
}

func (m *MinMaxService) IsFinished(gameBoard model.GameBoard) int {
	for i := range constants.FIELD_SIZE {
		cell1, _ := gameBoard.Cell(i, 0)
		cell2, _ := gameBoard.Cell(i, 1)
		cell3, _ := gameBoard.Cell(i, 2)
		if winner := checkLine(cell1, cell2, cell3); winner != constants.VOID {
			return winner
		}
		cell1, _ = gameBoard.Cell(0, i)
		cell2, _ = gameBoard.Cell(1, i)
		cell3, _ = gameBoard.Cell(2, i)
		if winner := checkLine(cell1, cell2, cell3); winner != constants.VOID {
			return winner
		}
	}
	cell1, _ := gameBoard.Cell(0, 0)
	cell2, _ := gameBoard.Cell(1, 1)
	cell3, _ := gameBoard.Cell(2, 2)
	if winner := checkLine(cell1, cell2, cell3); winner != constants.VOID {
		return winner
	}
	cell1, _ = gameBoard.Cell(0, 2)
	cell2, _ = gameBoard.Cell(1, 1)
	cell3, _ = gameBoard.Cell(2, 0)
	if winner := checkLine(cell1, cell2, cell3); winner != constants.VOID {
		return winner
	}

	for row := range constants.FIELD_SIZE {
		for col := range constants.FIELD_SIZE {
			if cell, _ := gameBoard.Cell(row, col); cell == constants.VOID {
				return constants.VOID
			}
		}
	}
	return constants.DRAW

}

func (m *MinMaxService) IsValid(gameBoard model.GameBoard, coordinate model.Coordinate) error {
	if cell, err := gameBoard.CellByCoordinate(coordinate); err != nil {
		return fmt.Errorf("cell validation error: %w", err)
	} else {
		if cell != constants.VOID {
			return fmt.Errorf("cell validation error: cell occupied")
		}
		return nil
	}

}

func checkLine(a, b, c int) int {
	if a != constants.VOID && a == b && a == c {
		return a
	}
	return constants.VOID
}

func getVoidCoordinates(gameBoard model.GameBoard) []model.Coordinate {
	result := make([]model.Coordinate, 0, constants.FIELD_SIZE*constants.FIELD_SIZE)
	for row := range constants.FIELD_SIZE {
		for col := range constants.FIELD_SIZE {
			if cell, _ := gameBoard.Cell(row, col); cell == constants.VOID {
				result = append(result, model.NewCoordinate(row, col))
			}
		}
	}
	return result
}

func (m *MinMaxService) minMax(gameBoard model.GameBoard, player int) Move {
	winner := m.IsFinished(gameBoard)
	switch winner {
	case constants.AI_PLAYER:
		return Move{coordinate: model.NewCoordinate(0, 0), score: 10}
	case constants.HUMAN_PLAYER:
		return Move{coordinate: model.NewCoordinate(0, 0), score: -10}
	case constants.DRAW:
		return Move{coordinate: model.NewCoordinate(0, 0), score: 0}
	}

	voidCoordinates := getVoidCoordinates(gameBoard)
	possibleMoves := make([]Move, 0, len(voidCoordinates))
	for _, coordinate := range voidCoordinates {
		nextGameBoard := gameBoard
		nextGameBoard.SetCell(coordinate, player)
		nextPlayer := constants.AI_PLAYER
		if player == constants.AI_PLAYER {
			nextPlayer = constants.HUMAN_PLAYER
		}
		nextMove := Move{coordinate: coordinate, score: m.minMax(nextGameBoard, nextPlayer).score}
		possibleMoves = append(possibleMoves, nextMove)
	}

	slices.SortFunc(possibleMoves, func(a, b Move) int {
		return a.score - b.score
	})
	if player == constants.AI_PLAYER {
		return possibleMoves[len(possibleMoves)-1]
	}
	return possibleMoves[0]

}

func (m *MinMaxService) SaveGame(game model.Game) {
	m.dataService.SaveGame(game)
}

func (m *MinMaxService) LoadGame(gameId uuid.UUID) (model.Game, error) {
	return m.dataService.LoadGame(gameId)
}

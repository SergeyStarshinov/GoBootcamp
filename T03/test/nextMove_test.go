package test

import (
	"testing"

	"tictactoe/internal/domain/model"
	"tictactoe/internal/infrastructure/constants"
)

func TestNextMoveDontLose(t *testing.T) {
	ai := constants.AI_PLAYER
	human := constants.HUMAN_PLAYER
	void := constants.VOID

	type testCase struct {
		input    model.GameBoard
		expected model.GameBoard
	}

	testCases := []testCase{
		{input: model.GameBoard{
			{void, void, ai},
			{human, void, void},
			{human, ai, void},
		},
			expected: model.GameBoard{
				{ai, void, ai},
				{human, void, void},
				{human, ai, void},
			},
		},
		{input: model.GameBoard{
			{ai, void, void},
			{human, human, void},
			{ai, void, human},
		},
			expected: model.GameBoard{
				{ai, void, void},
				{human, human, ai},
				{ai, void, human},
			},
		},
		{input: model.GameBoard{
			{ai, void, void},
			{human, human, ai},
			{human, ai, human},
		},
			expected: model.GameBoard{
				{ai, void, ai},
				{human, human, ai},
				{human, ai, human},
			},
		},
		{input: model.GameBoard{
			{void, void, void},
			{void, void, void},
			{human, void, void},
		},
			expected: model.GameBoard{
				{void, void, void},
				{void, ai, void},
				{human, void, void},
			},
		},
	}

	for _, testCase := range testCases {
		result := minMaxService.NextAiMove(testCase.input)
		if result != testCase.expected {
			t.Errorf("result is: %v, expected: %v.", result, testCase.expected)
		}
	}
}

func TestNextMoveWin(t *testing.T) {
	ai := constants.AI_PLAYER
	human := constants.HUMAN_PLAYER
	void := constants.VOID

	type testCase struct {
		input    model.GameBoard
		expected model.GameBoard
	}

	testCases := []testCase{
		{input: model.GameBoard{
			{ai, void, ai},
			{human, human, void},
			{human, ai, human},
		},
			expected: model.GameBoard{
				{ai, ai, ai},
				{human, human, void},
				{human, ai, human},
			},
		},
		{input: model.GameBoard{
			{ai, void, void},
			{human, ai, human},
			{ai, void, human},
		},
			expected: model.GameBoard{
				{ai, void, ai},
				{human, ai, human},
				{ai, void, human},
			},
		},
		{input: model.GameBoard{
			{ai, void, void},
			{human, ai, human},
			{human, ai, human},
		},
			expected: model.GameBoard{
				{ai, ai, void},
				{human, ai, human},
				{human, ai, human},
			},
		},
	}

	for _, testCase := range testCases {
		result := minMaxService.NextAiMove(testCase.input)
		if result != testCase.expected {
			t.Errorf("result is: %v, expected: %v.", result, testCase.expected)
		}
	}
}

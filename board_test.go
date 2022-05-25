package tictacgo

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	state := board.state
	if state != [9]int{} {
		t.Errorf("Unexpected state %q", state)
	}
}

func TestBoardFormatsCorrectly(t *testing.T) {
	tests := []struct {
		board    Board
		expected string
	}{
		{
			board:    NewBoard(),
			expected: "000000000",
		},
		{
			board:    Board{state: [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: "100000000",
		},
	}

	for _, test := range tests {
		result := test.board.String()
		if result != test.expected {
			t.Errorf("Expected %q, but got %q", test.expected, result)
		}
	}
}

func TestBoardMakeMove(t *testing.T) {
	tests := []struct {
		move     int
		board    Board
		expected [9]int
	}{
		{
			move:     5,
			board:    NewBoard(),
			expected: [9]int{0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			move:     5,
			board:    Board{state: [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: [9]int{1, 0, 0, 0, 0, 2, 0, 0, 0},
		},
	}

	for _, test := range tests {
		test.board.makeMove(test.move)
		if test.board.state != test.expected {
			t.Errorf("Expected %v but got %v", test.expected, test.board.state)
		}
	}

}

func TestBoardGetOutcome(t *testing.T) {
	tests := []struct {
		name     string
		board    Board
		expected GameOutcome
	}{
		{
			name:     "Blank board",
			board:    NewBoard(),
			expected: NotFinished,
		},
		{
			name:     "Draw",
			board:    Board{state: [9]int{1, 1, 2, 2, 2, 1, 1, 1, 2}},
			expected: Draw,
		},
		{
			name:     "Too many ones",
			board:    Board{state: [9]int{1, 1, 1, 2, 2, 1, 1, 1, 2}},
			expected: InvalidBoard,
		},
		{
			name:     "Horizontal 1 wins",
			board:    Board{state: [9]int{1, 1, 1, 2, 2, 1, 1, 2, 2}},
			expected: P1Wins,
		},
		{
			name:     "Horizontal 2 wins",
			board:    Board{state: [9]int{1, 1, 0, 2, 2, 2, 1, 1, 0}},
			expected: P2Wins,
		},
		{
			name:     "Vertical 1 wins",
			board:    Board{state: [9]int{1, 2, 2, 1, 2, 0, 1, 0, 0}},
			expected: P1Wins,
		},
		{
			name:     "Vertical 2 wins",
			board:    Board{state: [9]int{2, 2, 1, 1, 2, 1, 1, 2, 0}},
			expected: P2Wins,
		},
		{
			name:     "Diagonal 1 wins",
			board:    Board{state: [9]int{1, 2, 0, 0, 1, 2, 0, 0, 1}},
			expected: P1Wins,
		},
		{
			name:     "Diagonal 2 wins",
			board:    Board{state: [9]int{0, 1, 2, 1, 2, 0, 2, 1, 0}},
			expected: P2Wins,
		},
	}

	for _, test := range tests {
		outcome := test.board.getOutcome()
		if outcome != test.expected {
			t.Errorf("In scenario %q. Expected %q but got %q", test.name, test.expected, outcome)
		}
	}
}

func TestBoardGetValidMoves(t *testing.T) {
	tests := []struct {
		board    Board
		expected []int
	}{
		{
			board:    Board{state: [9]int{1, 1, 1, 2, 2, 0, 0, 0, 0}},
			expected: nil,
		},
		{
			board:    Board{state: [9]int{1, 1, 0, 2, 2, 0, 0, 0, 0}},
			expected: []int{2, 5, 6, 7, 8},
		},
		{
			board:    NewBoard(),
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, test := range tests {
		moves := test.board.getValidMoves()
		if test.expected == nil {
			if len(moves) > 0 {
				t.Errorf("Expected %v but got %v", test.expected, moves)
			}
		} else if !reflect.DeepEqual(moves, test.expected) {
			t.Errorf("Expected %v but got %v", test.expected, moves)
		}
	}
}

func TestBoardCopy(t *testing.T) {
	board := NewBoard()
	copied := board.copy()
	copied.makeMove(5)
	if reflect.DeepEqual(board.state, copied.state) {
		t.Error("Copy did not create an independent state")
	}
}

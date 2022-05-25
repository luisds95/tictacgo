package tictacgo

import (
	"fmt"
	"strings"
)

// GameOutcome is an enum of all possible board states
type GameOutcome int

// Each of the possible board states
const (
	Unknown GameOutcome = iota
	P1Wins
	P2Wins
	Draw
	NotFinished
	InvalidBoard
)

func (outcome GameOutcome) String() string {
	switch outcome {
	case P1Wins:
		return "Player 1 Wins"
	case P2Wins:
		return "Player 2 Wins"
	case Draw:
		return "Draw"
	case NotFinished:
		return "Game not yet finished"
	case InvalidBoard:
		return "Invalid board"
	}
	return "Unknown outcome"
}

func intToGameOutcome(value int) GameOutcome {
	switch value {
	case 1:
		return P1Wins
	case 2:
		return P2Wins
	default:
		return Unknown
	}
}

// NewBoard creates a blank board
func NewBoard() Board {
	return Board{state: [9]int{}}
}

// Board is the main container for a tictactoe board
type Board struct {
	state [9]int
}

func (board *Board) String() string {
	var strState [9]string
	for index, value := range board.state {
		strState[index] = fmt.Sprint(value)
	}
	return strings.Join(strState[:], "")
}

func (board *Board) makeMove(move int) {
	board.state[move] = board.nextPlayer()
}

func (board *Board) nextPlayer() int {
	var zeros int
	for _, value := range board.state {
		if value == 0 {
			zeros++
		}
	}

	if zeros%2 == 0 {
		return 2
	}
	return 1
}

func (board *Board) getOutcome() GameOutcome {
	if !board.isValid() {
		return InvalidBoard
	}

	// Find winner horizontally
	for i := 0; i < 9; i = i + 3 {
		if board.state[i] != 0 &&
			board.state[i] == board.state[i+1] &&
			board.state[i] == board.state[i+2] {
			return intToGameOutcome(board.state[i])
		}
	}

	// Find winner vertically
	for i := 0; i < 3; i++ {
		if board.state[i] != 0 &&
			board.state[i] == board.state[i+3] &&
			board.state[i] == board.state[i+6] {
			return intToGameOutcome(board.state[i])
		}
	}

	// Find winner diagonally
	if board.state[0] != 0 &&
		board.state[0] == board.state[4] &&
		board.state[0] == board.state[8] {
		return intToGameOutcome(board.state[0])
	}
	if board.state[2] != 0 &&
		board.state[2] == board.state[4] &&
		board.state[2] == board.state[6] {
		return intToGameOutcome(board.state[2])
	}

	// No winners
	if board.isFull() {
		return Draw
	}
	return NotFinished
}

func (board *Board) isValid() bool {
	var ones int
	var twos int
	for _, value := range board.state {
		if value == 1 {
			ones++
		} else if value == 2 {
			twos++
		}
	}
	return (ones == twos || ones == (twos+1))
}

func (board *Board) isFull() bool {
	for _, value := range board.state {
		if value == 0 {
			return false
		}
	}
	return true
}

func (board *Board) getValidMoves() []int {
	outcome := board.getOutcome()
	var validModes []int
	if outcome != NotFinished {
		return validModes
	}
	for index, value := range board.state {
		if value == 0 {
			validModes = append(validModes, index)
		}
	}
	return validModes
}

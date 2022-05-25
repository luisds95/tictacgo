package tictacgo

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Player interface {
	getAction(board Board) int
}

type HumanPlayer struct {
	number int
}

func NewHumanPlayer(number int) HumanPlayer {
	return HumanPlayer{number: number}
}

func (player HumanPlayer) getAction(board Board) int {
	validActions := board.getValidMoves()
	for {
		var input int
		fmt.Print("Insert your next move:")
		fmt.Scanf("%d", &input)
		if slices.Contains(validActions, input) {
			return input
		}
	}
}

type ExhaustiveSearchPlayer struct {
	number     int
	winOutcome GameOutcome
	rewardWin  float64
	rewardDraw float64
	rewardLose float64
	database   Database
}

func NewExhaustiveSearchPlayer(number int, database Database) ExhaustiveSearchPlayer {
	winOutcome := intToGameOutcome(number)
	return ExhaustiveSearchPlayer{
		number:     number,
		winOutcome: winOutcome,
		rewardWin:  1,
		rewardDraw: 0,
		rewardLose: -1,
		database:   database,
	}
}

func (player ExhaustiveSearchPlayer) getAction(board Board) int {
	values := player.database.get(board)
	if values == nil {
		values = player.evaluateMoves(board)
	}

	if board.nextPlayer() == 1 {
		return getMaxActionValue(values)
	} else {
		return getMinActionValue(values)
	}
}

func (player ExhaustiveSearchPlayer) evaluateMoves(board Board) map[string]float64 {
	values := map[string]float64{}
	for _, move := range board.getValidMoves() {
		newBoard := board.copy()
		newBoard.makeMove(move)
		outcome := newBoard.getOutcome()
		strMove := fmt.Sprint(move)
		switch outcome {
		case NotFinished:
			fmt.Println(NotFinished)
		case Draw:
			values[strMove] = player.rewardDraw
		case player.winOutcome:
			values[strMove] = player.rewardWin
		default:
			values[strMove] = player.rewardLose
		}
	}
}

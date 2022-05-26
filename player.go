package tictacgo

import (
	"fmt"
	"time"

	"strconv"

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
	number          int
	winOutcome      GameOutcome
	rewardWin       float64
	rewardDraw      float64
	rewardLose      float64
	database        Database
	isTraining      bool
	commitFrequency int
}

func NewExhaustiveSearchPlayer(number int, database Database) ExhaustiveSearchPlayer {
	winOutcome := intToGameOutcome(number)
	return ExhaustiveSearchPlayer{
		number:          number,
		winOutcome:      winOutcome,
		rewardWin:       1,
		rewardDraw:      0,
		rewardLose:      -1,
		database:        database,
		isTraining:      false,
		commitFrequency: 1000,
	}
}

func (player *ExhaustiveSearchPlayer) train() {
	initialTime := time.Now()
	fmt.Println("Starting training at", initialTime)

	player.isTraining = true

	board := NewBoard()
	player.evaluateMoves(board)
	player.database.commit()

	player.isTraining = false

	finalTime := time.Now()
	fmt.Println("Training finished! It took", finalTime.Sub(initialTime))
}

func (player ExhaustiveSearchPlayer) getAction(board Board) int {
	values := player.database.get(board)
	if values == nil {
		values = player.evaluateMoves(board)
	}

	shouldMax := board.nextPlayer() == player.number
	bestAction, _ := getBestActionFromValues(values, shouldMax)

	return bestAction
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
			innerValues := player.evaluateMoves(newBoard)
			shouldMax := newBoard.nextPlayer() == player.number
			_, bestValue := getBestActionFromValues(innerValues, shouldMax)
			values[strMove] = bestValue
		case Draw:
			values[strMove] = player.rewardDraw
		case player.winOutcome:
			values[strMove] = player.rewardWin
		default:
			values[strMove] = player.rewardLose
		}
	}

	if player.isTraining {
		player.database.update(board, values)
		if player.database.size()%player.commitFrequency == 0 {
			player.database.commit()
		}
	}

	return values
}

func getBestActionFromValues(values map[string]float64, max bool) (int, float64) {
	bestValue := 100.0
	if max {
		bestValue *= -1
	}
	bestAction := -1
	for actionStr, value := range values {
		action, err := strconv.Atoi(actionStr)
		if err != nil {
			panic(fmt.Sprintf("Unexpected error: %v", err))
		}
		if (max && value > bestValue) || (!max && value < bestValue) {
			bestValue = value
			bestAction = action
		}
	}
	return bestAction, bestValue
}

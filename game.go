package tictacgo

import "fmt"

func Play() {
	players := map[int]Player{
		1: NewHumanPlayer(),
		2: NewHumanPlayer(),
	}
	board := NewBoard()

	for board.getOutcome() == NotFinished {
		fmt.Println(board.Pretty())
		action := players[board.nextPlayer()].getAction(board)
		board.makeMove(action)
	}
	fmt.Println(board.Pretty())
	fmt.Println(board.getOutcome())
}

package tictacgo

import "fmt"

func getDefaultDbName() string {
	return "stateValues.json"
}

func Play() {
	database := NewMapDB(getDefaultDbName())
	players := map[int]Player{
		1: NewHumanPlayer(1),
		2: NewExhaustiveSearchPlayer(2, database),
	}
	board := NewBoard()

	for board.getOutcome() == NotFinished {
		fmt.Println("\n" + board.Pretty())
		action := players[board.nextPlayer()].getAction(board)
		board.makeMove(action)
	}
	fmt.Println(board.Pretty())
	fmt.Println(board.getOutcome())
}

func Train() {
	database := NewMapDB(getDefaultDbName())
	player := NewExhaustiveSearchPlayer(1, database)
	player.train()
}

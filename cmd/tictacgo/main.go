package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/luisds95/tictacgo"
)

func main() {
	args := os.Args
	executionMode := strings.ToLower(args[1])
	if executionMode == "train" {
		tictacgo.Train()
	} else if executionMode == "play" {
		tictacgo.Play()
	} else {
		fmt.Println("Unknown execution mode! Try with one of train, play")
	}
}

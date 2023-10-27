package main

import (
	"fmt"

	playground "github.com/Friedfisch/tictactoe/playGround"
)

func main() {
	var size = 5
	f := playground.NewPlayGround(size, 2)
	fmt.Printf("Board: %d\n", f.Board())
	fmt.Printf("Players: %d\n", f.Players())
	for i := byte(1); i < f.Players(); i++ {
		fmt.Printf("Events for player %d: %d\n", i, f.Log(i))
	}
}

// bowling_game_app project main.go
package main

import (
	"app/bowling_game"
	"fmt"
)

func main() {
	fmt.Println("bowling_game_app.exe")

	// replicate logic of the final test
	g := &bowling_game.Game{}
	for i := 0; i < 12; i++ {
		g.Roll(10)
	}
	fmt.Println(g.Score())
}

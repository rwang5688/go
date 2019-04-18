// package must be "main" to install $(directory_name).exe to $(GOPATH)/bin
package main

import (
	"app/bowling_game"
	"fmt"
)

func main() {
	fmt.Println("bowling_game_app.exe")
	g := &bowling_game.Game{}
	for i := 0; i < 12; i++ {
		g.Roll(10)
	}
	fmt.Println(g.Score())
}

// bowling_game_unittest project game_test.go
package bowling_game_unittest

import (
	"app/bowling_game"
	"strconv"
	"testing"
)

// Setup
func setUp() *bowling_game.Game {
	g := &bowling_game.Game{}
	g.Init()
	return g
}

// 1st Test
func TestGutterGame(t *testing.T) {
	// Arrange
	g := setUp()

	// Act
	rollMany(g, 20, 0)

	// Assert
	s := g.Score()
	if s != 0 {
		t.Error(`Expected score is 0; score is ` + strconv.Itoa(s))
	}
}

// 2nd Test
func TestAllOnes(t *testing.T) {
	// Arrange
	g := setUp()

	// Act
	rollMany(g, 20, 1)

	// Assert
	s := g.Score()
	if s != 20 {
		t.Error(`Expected score is 20; score is ` + strconv.Itoa(s))
	}
}

// 3rd Test
func TestOneSpare(t *testing.T) {
	// Arrange
	g := setUp()

	// Act
	rollSpare(g)
	g.Roll(1) // spare bonus = 1
	rollMany(g, 17, 0)

	// Assert
	s := g.Score()
	if s != 12 {
		t.Error(`Expected score is 12; score is ` + strconv.Itoa(s))
	}
}

// 4th Test
func TestOneStrike(t *testing.T) {
	// Arrange
	g := setUp()

	// Act
	g.Roll(10) // strike
	g.Roll(1)
	g.Roll(1) // strike bonus = 2
	rollMany(g, 16, 0)

	// Assert
	s := g.Score()
	if s != 14 {
		t.Error(`Expected score is 14; score is ` + strconv.Itoa(s))
	}
}

// 5th Test
func TestPerfectGame(t *testing.T) {
	// Arrange
	g := setUp()

	// Act
	rollMany(g, 12, 10) // 12 strikes

	// Assert
	s := g.Score()
	if s != 300 {
		t.Error(`Expected score is 300; score is ` + strconv.Itoa(s))
	}
}

// game test helper functions
func rollMany(g *bowling_game.Game, n int, pins int) {
	for i := 0; i < n; i++ {
		g.Roll(pins)
	}
}

func rollSpare(g *bowling_game.Game) {
	g.Roll(5)
	g.Roll(5)
}

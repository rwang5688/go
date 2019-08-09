// bowling_game project game.go
package bowling_game

// separate responsibilty for keeping track and calculating score
type Game struct {
	pins_        [21]int
	currentRoll_ int
}

// initialize struct variables
func (g *Game) Init() {
	for i := 0; i < 21; i++ {
		g.pins_[i] = 0
	}
	g.currentRoll_ = 0
}

// keep track of rolls
func (g *Game) Roll(pins int) {
	g.pins_[g.currentRoll_] = pins
	g.currentRoll_++
}

// calculate score
func (g Game) Score() int {
	score := 0

	rollIndex := 0
	for frameIndex := 0; frameIndex < 10; frameIndex++ {
		if g.isStrike(rollIndex) {
			score = score + 10 + g.strikeBonus(rollIndex)
			rollIndex += 1
		} else if g.isSpare(rollIndex) {
			score = score + 10 + g.spareBonus(rollIndex)
			rollIndex += 2
		} else {
			score += g.sumOfPinsInFrame(rollIndex)
			rollIndex += 2
		}
	}

	return score
}

// game rule helper functions
func (g *Game) isStrike(rollIndex int) bool {
	return g.pins_[rollIndex] == 10
}

func (g *Game) strikeBonus(rollIndex int) int {
	return g.pins_[rollIndex+1] + g.pins_[rollIndex+2]
}

func (g *Game) isSpare(rollIndex int) bool {
	return g.pins_[rollIndex]+g.pins_[rollIndex+1] == 10
}

func (g *Game) spareBonus(rollIndex int) int {
	return g.pins_[rollIndex+2]
}

func (g *Game) sumOfPinsInFrame(rollIndex int) int {
	return g.pins_[rollIndex] + g.pins_[rollIndex+1]
}

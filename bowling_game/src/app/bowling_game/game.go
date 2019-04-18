package bowling_game

type Game struct {
	rolls_ [21]int
	currentRoll_ int
}

// initialize struct variables
func (g *Game) Init() {
	for i:=0; i < 21; i++ {
		g.rolls_[i] = 0
	} 
	g.currentRoll_ = 0
}

// keep track of rolls
func (g *Game) Roll(pins int) {
	g.rolls_[g.currentRoll_] = pins
	g.currentRoll_++
}

// calculate score
func (g Game) Score() int {
	score := 0
	rollIndex := 0
	for frame := 0; frame < 10; frame++ {
		if g.isStrike(rollIndex) {
			score += 10 + g.strikeBonus(rollIndex)
			rollIndex += 1
		} else if g.isSpare(rollIndex) {
			score += 10 + g.spareBonus(rollIndex)
			rollIndex += 2
		} else { // open frame
			score += g.sumOfBallsInFrame(rollIndex)
			rollIndex += 2
		}
	}

	return score
}

// game rule helper functions
func (g Game) isStrike(rollIndex int) bool {
	return g.rolls_[rollIndex] == 10
} 

func (g Game) strikeBonus(rollIndex int) int {
	return g.rolls_[rollIndex+1] + g.rolls_[rollIndex+2]
}

func (g Game) isSpare(rollIndex int) bool {
	return g.rolls_[rollIndex] + g.rolls_[rollIndex+1] == 10
} 

func (g Game) spareBonus(rollIndex int) int {
	return g.rolls_[rollIndex+2]
}

func (g Game) sumOfBallsInFrame(rollIndex int) int {
	return g.rolls_[rollIndex] + g.rolls_[rollIndex+1]
}

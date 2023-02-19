package eval

import "math"

const (
	// default K-factor
	K = 32

	// default deviation
	D = 400
)

type Elo struct {
	K int
	D int
}

func NewDefaultElo() *Elo {
	return &Elo{K, D}
}

func (e *Elo) ExpectedWinPercentage(rA, rB int) float64 {
	return 1 / (1 + math.Pow(10, float64(rB-rA)/float64(D)))
}

func (e *Elo) Rating(rA, rB int, won bool) int {
	score := 0.0
	if won {
		score = 1.0
	}

	eA := e.ExpectedWinPercentage(rA, rB)
	return rA + (int(K * (score - eA)))
}

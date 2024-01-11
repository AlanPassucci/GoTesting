package prey

import "math/rand"

type tuna struct {
	// max speed in m/s
	maxSpeed float64
}

func NewTuna() Prey {
	return &tuna{
		maxSpeed: 252,
	}
}

func (t *tuna) GetSpeed() float64 {
	return t.maxSpeed * rand.Float64()
}

// StubTuna is a stub implementation of the Prey interface for testing
type StubTuna struct {
	SpeedOnGetSpeed float64
}

// GetSpeed returns the speed on GetSpeed
func (st *StubTuna) GetSpeed() float64 {
	return st.SpeedOnGetSpeed
}

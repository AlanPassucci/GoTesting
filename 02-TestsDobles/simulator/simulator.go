package simulator

import (
	"fmt"
	"math/big"
)

type CatchSimulator interface {
	CanCatch(distance float64, speed float64, catchSpeed float64) bool
	GetLinearDistance(position [2]float64) float64
}

type catchSimulator struct {
	// max time to catch the prey in seconds
	maxTimeToCatch float64
}

func NewCatchSimulator(maxTimeToCatch float64) CatchSimulator {
	return &catchSimulator{
		maxTimeToCatch: maxTimeToCatch,
	}
}

func (r *catchSimulator) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	fmt.Printf("Distance: (%.2f) | Prey Speed: %.2f m/s\n", distance, catchSpeed)
	timeToCatch := distance / (speed - catchSpeed)
	fmt.Printf("Time to catch: %.2f\n seconds", timeToCatch)
	return timeToCatch > 0 && timeToCatch <= r.maxTimeToCatch
}

func (r *catchSimulator) GetLinearDistance(position [2]float64) float64 {
	x := big.NewFloat(position[0])
	y := big.NewFloat(position[1])
	z := x.Add(x.Mul(x, x), y.Mul(y, y))
	res, _ := z.Sqrt(z).Float64()
	fmt.Printf("Distance: %.2f meters\n", res)
	return res
}

// SpySimulator is a Spy implementation of the CatchSimulator interface for testing
type SpySimulator struct {
	GetLinearDistanceCalledTimes int
	CaughtPreySpeed              float64
}

// MockSimulator is a Mock implementation of the CatchSimulator interface for testing
type MockSimulator struct {
	ResultOnCanCatch                  bool
	LinearDistanceOnGetLinearDistance float64
	SpySimulator                      *SpySimulator
}

func (m *MockSimulator) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	m.SpySimulator.GetLinearDistanceCalledTimes++
	m.SpySimulator.CaughtPreySpeed = catchSpeed
	return m.ResultOnCanCatch
}

func (m *MockSimulator) GetLinearDistance(position [2]float64) float64 {
	return m.LinearDistanceOnGetLinearDistance
}

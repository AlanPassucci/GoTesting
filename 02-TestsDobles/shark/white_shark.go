package shark

import (
	"fmt"
	"gotesting/02-TestsDobles/prey"
	"gotesting/02-TestsDobles/simulator"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type whiteShark struct {
	// speed in m/s
	speed float64
	// position of the shark in the map of 500 * 500 meters
	position [2]float64
	// simulator
	simulator simulator.CatchSimulator
}

func NewWhiteShark(simulator simulator.CatchSimulator) Shark {
	return &whiteShark{
		speed:     144,
		position:  getInitialPosition(),
		simulator: simulator,
	}
}

func (w *whiteShark) Hunt(prey prey.Prey) error {
	if w.simulator.CanCatch(w.simulator.GetLinearDistance(w.position), w.speed, prey.GetSpeed()) {
		fmt.Println("ñam ñam")
		return nil
	}
	return fmt.Errorf("could not hunt the prey")
}

func getInitialPosition() [2]float64 {
	x := rand.Float64() * 500
	rand.Seed(time.Now().UnixNano())
	y := rand.Float64() * 500
	return [2]float64{x, y}
}

package shark

import (
	"fmt"
	"gotesting/03-TestsIntegracion/pkg/storage"
	"gotesting/03-TestsIntegracion/prey"
	"gotesting/03-TestsIntegracion/simulator"
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

func NewWhiteShark(simulator simulator.CatchSimulator, storage storage.Storage) Shark {
	return &whiteShark{
		speed:     storage.GetValue("white_shark_speed").(float64),
		position:  [2]float64{storage.GetValue("white_shark_x").(float64), storage.GetValue("white_shark_y").(float64)},
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

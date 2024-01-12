package shark

import (
	"gotesting/04-TestsFuncionales/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) (error, float64)
	Configure(position [2]float64, speed float64)
}

package shark

import (
	"gotesting/03-TestsIntegracion/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) error
}

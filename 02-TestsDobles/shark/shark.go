package shark

import (
	"gotesting/02-TestsDobles/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) error
}

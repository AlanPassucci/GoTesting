package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatchSimulator_CanCatch(t *testing.T) {
	t.Run("it should return true if the time to catch is greater than 0 and less than or equal to the max time to catch", func(t *testing.T) {
		// Arrange
		distance := 15.0
		speed := 4.8
		catchSpeed := 2.75
		mockSimulator := &MockSimulator{
			ResultOnCanCatch:                  true,
			LinearDistanceOnGetLinearDistance: distance,
			SpySimulator:                      &SpySimulator{},
		}

		// Act
		obtainedResult := mockSimulator.CanCatch(distance, speed, catchSpeed)

		// Assert
		assert.True(t, obtainedResult)
		assert.Equal(t, 1, mockSimulator.SpySimulator.GetLinearDistanceCalledTimes)
	})
}

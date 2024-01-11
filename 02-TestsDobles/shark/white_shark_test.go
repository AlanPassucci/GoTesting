package shark

import (
	"gotesting/02-TestsDobles/prey"
	"gotesting/02-TestsDobles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhiteShark_Hunt(t *testing.T) {
	t.Run("it should return nil if the shark can catch the prey and GetLinearDistance was called", func(t *testing.T) {
		// Arrange
		mockSimulator := &simulator.MockSimulator{
			ResultOnCanCatch:                  true,
			LinearDistanceOnGetLinearDistance: 15.0,
			SpySimulator:                      &simulator.SpySimulator{},
		}
		whiteShark := NewWhiteShark(mockSimulator)
		stubTuna := &prey.StubTuna{
			SpeedOnGetSpeed: 100.0,
		}

		// Act
		err := whiteShark.Hunt(stubTuna)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 1, mockSimulator.SpySimulator.GetLinearDistanceCalledTimes)
		assert.Equal(t, stubTuna.GetSpeed(), mockSimulator.SpySimulator.CaughtPreySpeed)
	})

	t.Run("it should return an error when the shark is slower than the prey", func(t *testing.T) {
		// Arrange
		mockSimulator := &simulator.MockSimulator{
			ResultOnCanCatch:                  false,
			LinearDistanceOnGetLinearDistance: 15.0,
			SpySimulator:                      &simulator.SpySimulator{},
		}
		whiteShark := NewWhiteShark(mockSimulator)
		stubTuna := &prey.StubTuna{
			SpeedOnGetSpeed: 200.0, // Shark is slower than the prey
		}
		expectedError := "could not hunt the prey"

		// Act
		err := whiteShark.Hunt(stubTuna)

		// Assert
		assert.EqualError(t, err, expectedError)
		assert.Equal(t, 1, mockSimulator.SpySimulator.GetLinearDistanceCalledTimes)
		assert.Equal(t, stubTuna.GetSpeed(), mockSimulator.SpySimulator.CaughtPreySpeed)
	})

	t.Run("it should return an error when the shark is fast but the distance is too large", func(t *testing.T) {
		// Arrange
		mockSimulator := &simulator.MockSimulator{
			ResultOnCanCatch:                  false,
			LinearDistanceOnGetLinearDistance: 150.0,
			SpySimulator:                      &simulator.SpySimulator{},
		}
		whiteShark := NewWhiteShark(mockSimulator)
		stubTuna := &prey.StubTuna{
			SpeedOnGetSpeed: 100.0, // Shark is faster than the prey
		}
		expectedError := "could not hunt the prey"

		// Act
		err := whiteShark.Hunt(stubTuna)

		// Assert
		assert.EqualError(t, err, expectedError)
		assert.Equal(t, 1, mockSimulator.SpySimulator.GetLinearDistanceCalledTimes)
		assert.Equal(t, stubTuna.GetSpeed(), mockSimulator.SpySimulator.CaughtPreySpeed)
	})
}

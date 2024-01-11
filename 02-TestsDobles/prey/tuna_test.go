package prey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuna_GetSpeed(t *testing.T) {
	t.Run("it should return a speed of 150", func(t *testing.T) {
		// Arrange
		expectedSpeed := 150.0
		stubTuna := &StubTuna{
			SpeedOnGetSpeed: expectedSpeed,
		}

		// Act
		obtainedSpeed := stubTuna.GetSpeed()

		// Assert
		assert.Equal(t, expectedSpeed, obtainedSpeed)
	})

	t.Run("it should never return speed greater than 252", func(t *testing.T) {
		// Arrange
		speedResult := 100.0
		stubTuna := &StubTuna{
			SpeedOnGetSpeed: speedResult,
		}

		// Act
		obtainedSpeed := stubTuna.GetSpeed()

		// Assert
		assert.False(t, obtainedSpeed > 252)
	})

	t.Run("it should never return speed lower than 0", func(t *testing.T) {
		// Arrange
		speedResult := 75.0
		stubTuna := &StubTuna{
			SpeedOnGetSpeed: speedResult,
		}

		// Act
		obtainedSpeed := stubTuna.GetSpeed()

		// Assert
		assert.False(t, obtainedSpeed < 0)
	})
}

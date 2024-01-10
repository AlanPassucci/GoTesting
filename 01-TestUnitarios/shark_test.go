package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	// Arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  10,
	}
	prey := Prey{
		name:  "fish",
		speed: 9,
	}

	// Act
	err := shark.Hunt(&prey)

	// Assert
	assert.NoError(t, err)
	assert.False(t, shark.hungry)
	assert.True(t, shark.tired)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// Arrange
	shark := Shark{
		hungry: true,
		tired:  true,
		speed:  10,
	}
	prey := Prey{
		name:  "fish",
		speed: 9,
	}

	// Act
	err := shark.Hunt(&prey)

	// Assert
	assert.True(t, shark.tired)
	assert.Error(t, err)
	assert.EqualError(t, err, "cannot hunt, i am really tired")
}

func TestSharkCannotHuntBecauseIsNotHungry(t *testing.T) {
	// Arrange
	shark := Shark{
		hungry: false,
		tired:  false,
		speed:  10,
	}
	prey := Prey{
		name:  "fish",
		speed: 9,
	}

	// Act
	err := shark.Hunt(&prey)

	// Assert
	assert.False(t, shark.hungry)
	assert.Error(t, err)
	assert.EqualError(t, err, "cannot hunt, i am not hungry")
}

func TestSharkCannotReachThePrey(t *testing.T) {
	// Arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  10,
	}
	prey := Prey{
		name:  "fish",
		speed: 11,
	}

	// Act
	err := shark.Hunt(&prey)

	// Assert
	assert.True(t, shark.tired)
	assert.Error(t, err)
	assert.EqualError(t, err, "could not catch it")
}

func TestSharkHuntNilPrey(t *testing.T) {
	// Arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  10,
	}
	var prey *Prey

	// Act
	err := shark.Hunt(prey)

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "prey is nil")
}

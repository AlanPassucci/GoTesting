package shark

import (
	"gotesting/03-TestsIntegracion/pkg/storage"
	"gotesting/03-TestsIntegracion/prey"
	"gotesting/03-TestsIntegracion/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhiteShark_Hunt(t *testing.T) {
	t.Run("should return nil when the shark can catch the prey", func(t *testing.T) {
		// Arrange
		mockStorage := storage.MockStorage{}

		mockStorage.On("GetValue", "tuna_speed").Return(100.0)
		mockStorage.On("GetValue", "white_shark_speed").Return(150.0)
		mockStorage.On("GetValue", "white_shark_x").Return(250.0)
		mockStorage.On("GetValue", "white_shark_y").Return(250.0)
		mockStorage.On("GetLinearDistance").Return(50.0).Maybe()

		tuna := prey.NewTuna(&mockStorage)
		simulator := simulator.NewCatchSimulator(10)
		whiteShark := NewWhiteShark(simulator, &mockStorage)

		// Act
		err := whiteShark.Hunt(tuna)

		// Assert
		assert.Nil(t, err)
		assert.True(t, mockStorage.AssertExpectations(t))
	})

	t.Run("should return error because the shark has slower speed", func(t *testing.T) {
		// Arrange
		expectedError := "could not hunt the prey"
		mockStorage := storage.MockStorage{}

		mockStorage.On("GetValue", "tuna_speed").Return(150.0)
		mockStorage.On("GetValue", "white_shark_speed").Return(100.0)
		mockStorage.On("GetValue", "white_shark_x").Return(250.0)
		mockStorage.On("GetValue", "white_shark_y").Return(250.0)
		mockStorage.On("GetLinearDistance").Return(50.0).Maybe()

		tuna := prey.NewTuna(&mockStorage)
		simulator := simulator.NewCatchSimulator(10)
		whiteShark := NewWhiteShark(simulator, &mockStorage)

		// Act
		err := whiteShark.Hunt(tuna)

		// Assert
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedError)
		assert.True(t, mockStorage.AssertExpectations(t))
	})

	t.Run("should return error because the shark has fastest speed but the distance is longer", func(t *testing.T) {
		// Arrange
		expectedError := "could not hunt the prey"
		mockStorage := storage.MockStorage{}

		mockStorage.On("GetValue", "tuna_speed").Return(100.0)
		mockStorage.On("GetValue", "white_shark_speed").Return(150.0)
		mockStorage.On("GetValue", "white_shark_x").Return(250.0)
		mockStorage.On("GetValue", "white_shark_y").Return(250.0)
		mockStorage.On("GetLinearDistance").Return(2000.0).Maybe()

		tuna := prey.NewTuna(&mockStorage)
		simulator := simulator.NewCatchSimulator(10)
		whiteShark := NewWhiteShark(simulator, &mockStorage)

		// Act
		err := whiteShark.Hunt(tuna)

		// Assert
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedError)
		assert.True(t, mockStorage.AssertExpectations(t))
	})
}

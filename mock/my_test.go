package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// In your Go test file
func TestMyFunction(t *testing.T) {
	// Test case 1: Positive inputs
	sum, err := MyFunction(2, 7)
	assert.Equal(t, 5, sum, "Sum of 2 and 3 is 5")
	assert.NoError(t, err, "There are no errors")

}

// In your Go test file
func TestMyFunction2(t *testing.T) {
	// Test case 2: One negative input
	sum, err := MyFunction(-2, 3)
	assert.Equal(t, 0, sum, "Sum of -2 and 3 is 0")
	assert.Error(t, err, "Error is not nil for negative input")

}

///-------------------------------
///-------------------------------
///-------------------------------

type mockCalculateArea struct {
	mock.Mock
}

func (m *mockCalculateArea) calculateArea(width int, height int) int {
	args := m.Called(width, height)
	return args.Int(0)
}

// In your Go test file
func TestCalculateArea(t *testing.T) {
	// Create a mock object for the calculateArea function
	mockObj := new(mockCalculateArea)

	// Set up the expected return value
	mockObj.On("calculateArea", 5, 10).Return(50)

	// Call the function and check the result
	actualArea := mockObj.calculateArea(5, 10)
	assert.Equal(t, 50, actualArea, "The calculated area is incorrect")

	// Verify that the expected function call was made
	mockObj.AssertExpectations(t)
}

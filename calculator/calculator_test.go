package calculator_test

import (
	"testing"

	"github.com/burubur/fastcampus-golang/calculator"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(1, 2)
	expected := 3

	assert.Equal(t, expected, result, "Add(1, 2) should return 3")
}

func TestDivideBy(t *testing.T) {
	result, err := calculator.DivideBy(10, 2)
	assert.NoError(t, err, "it should not return error")

	assert.Equal(t, 5, result, "DivideBy(10, 2) should return 5")
}

func TestDivideBy_BiggerNumber(t *testing.T) {
	result, err := calculator.DivideBy(2, 10)
	assert.Error(t, err, "it should return error")

	assert.Equal(t, 0, result, "DivideBy(2, 10) should return 0 due to invalid number")
}

func TestDivideBy_ZeroNumber(t *testing.T) {
	assert.NotPanics(t, func() {
		_, _ = calculator.DivideBy(0, 10)
	}, "It should not panic even if we send an invalid number, at least it should return 0")
}

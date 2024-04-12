package calculator_test

import (
	"testing"

	"github.com/burubur/fastcampus-golang/calculator"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) should return 3 but got %d", result)
	}
}

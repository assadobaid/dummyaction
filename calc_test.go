package calc

import (
	"fmt"
	"testing"
)

// Add returns the sum of two integers
func TestAdd(t *testing.T) {
	fmt.Printf("Hej Hej")
	if Add(6, 3) != 9 {
		t.Fail()
	}
}

// Subtract returns the subtraction of two integers
func TestSubtract(t *testing.T) {
	if Sub(6, 3) != 3 {
		t.Fail()
	}
}

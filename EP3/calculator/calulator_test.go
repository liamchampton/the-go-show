package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 5
	got := Add(2, 3)
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

func TestAddEdgeCase(t *testing.T) {
	want := math.MaxInt64
	got := Add(math.MaxInt64-1, 1) // -1 prevents integer overflow
	if got != want {
		t.Errorf("Add(%d, %d) = %d; want %d", math.MaxInt64-1, 1, got, want)
	}
}

func TestSubtract(t *testing.T) {
	want := 2
	got := Subtract(5, 3)
	if got != want {
		t.Errorf("Subtract(5, 3) = %d; want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	want := 6
	got := Multiply(2, 3)
	if got != want {
		t.Errorf("Multiply(2, 3) = %d; want %d", got, want)
	}
}

func TestDivide(t *testing.T) {
	want := 2.0
	got := Divide(6, 3)
	if got != want {
		t.Errorf("Divide(6, 3) = %f; want %f", got, want)
	}
}

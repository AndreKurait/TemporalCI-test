package main

import "testing"

func TestClamp(t *testing.T) {
	tests := []struct {
		val, min, max, want int
	}{
		{5, 0, 10, 5},
		{-1, 0, 10, 0},
		{15, 0, 10, 10},
		{0, 0, 10, 0},
		{10, 0, 10, 10},
	}
	for _, tt := range tests {
		if got := Clamp(tt.val, tt.min, tt.max); got != tt.want {
			t.Errorf("Clamp(%d, %d, %d) = %d, want %d", tt.val, tt.min, tt.max, got, tt.want)
		}
	}
}

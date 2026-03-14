package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		input, want int
	}{
		{5, 5},
		{-3, 3},
		{0, 0},
		{-100, 100},
	}
	for _, tt := range tests {
		if got := Abs(tt.input); got != tt.want {
			t.Errorf("Abs(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

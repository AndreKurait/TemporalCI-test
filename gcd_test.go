package main

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{12, 8, 4}, {7, 13, 1}, {0, 5, 5}, {100, 75, 25}, {-6, 9, 3},
	}
	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

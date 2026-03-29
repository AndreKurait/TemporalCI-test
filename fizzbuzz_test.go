package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1, "1"}, {2, "2"}, {3, "Fizz"}, {4, "4"}, {5, "Buzz"},
		{6, "Fizz"}, {10, "Buzz"}, {15, "FizzBuzz"}, {30, "FizzBuzz"},
		{7, "7"}, {9, "Fizz"}, {25, "Buzz"}, {45, "FizzBuzz"},
	}
	for _, tt := range tests {
		if got := FizzBuzz(tt.input); got != tt.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}


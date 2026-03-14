package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"abba", true},
		{"abc", false},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.input); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

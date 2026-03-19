package main

import "testing"

func TestAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"listen", "silent", true},
		{"hello", "world", false},
		{"Astronomer", "Moon starer", false},
		{"abc", "cba", true},
		{"", "", true},
		{"a", "ab", false},
	}
	for _, tt := range tests {
		if got := Anagram(tt.a, tt.b); got != tt.want {
			t.Errorf("Anagram(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

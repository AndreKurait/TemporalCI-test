package main

import "testing"

func TestTruncate(t *testing.T) {
	tests := []struct {
		input  string
		maxLen int
		want   string
	}{
		{"hello", 10, "hello"},
		{"hello world", 5, "he..."},
		{"hi", 2, "hi"},
		{"abcdef", 3, "abc"},
	}
	for _, tt := range tests {
		if got := Truncate(tt.input, tt.maxLen); got != tt.want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tt.input, tt.maxLen, got, tt.want)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"hello world", 2},
		{"", 0},
		{"  spaces  everywhere  ", 2},
		{"one", 1},
	}
	for _, tt := range tests {
		if got := CountWords(tt.input); got != tt.want {
			t.Errorf("CountWords(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

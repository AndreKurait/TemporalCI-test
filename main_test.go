package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("TemporalCI")
	want := "Hello, TemporalCI!"
	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}

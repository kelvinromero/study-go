package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	// Creaetes a Buffer, gets it's address
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	// Calls the function
	Countdown(buffer, spySleeper)

	// Saves what the buffer captured
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleep, want 3, got %d", spySleeper.Calls)
	}
}

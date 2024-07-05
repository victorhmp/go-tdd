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
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	expected := `3
2
1
Go!`

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("Not enough calls to sleeper. Expected 3 got %d", spySleeper.Calls)
	}
}

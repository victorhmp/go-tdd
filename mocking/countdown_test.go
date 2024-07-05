package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleepOperation = "sleep"
const writeOperation = "write"

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleepOperation)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, writeOperation)

	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("prints a countdown from 3, ending in 'Go!'", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleepWrite := &SpyCountdownOperations{}

		Countdown(buffer, spySleepWrite)

		got := buffer.String()
		expected := `3
2
1
Go!`

		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("sleep is called before every print", func(t *testing.T) {
		spyOperations := &SpyCountdownOperations{}

		// Since our SpyCountdownOperations implements both io.Writer and Sleeper,
		// we can pass it to both arguments Countdown expects.
		Countdown(spyOperations, spyOperations)

		expected := []string{
			writeOperation,
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
		}

		if !reflect.DeepEqual(expected, spyOperations.Calls) {
			t.Errorf("Expected calls %v got %v", expected, spyOperations.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

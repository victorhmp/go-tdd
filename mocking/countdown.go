package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(output io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(output, i)
		sleeper.Sleep()
	}
	fmt.Fprint(output, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

// This is a nice example where using dependency ejection and mocking saves
// us from havin tests that take 3 seconds to complete.

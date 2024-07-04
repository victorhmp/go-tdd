package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(output io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(output, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(output, finalWord)
}

func main() {
	Countdown(os.Stdout)
}

// This is a nice example where using dependency ejection and mocking saves
// us from havin tests that take 3 seconds to complete.

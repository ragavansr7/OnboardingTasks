package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// func Countdown(out io.Writer) {
// 	fmt.Fprint(out, "3")
// }

const finalWord = "Go!"
const countdownStart = 3

// func Countdown(out io.Writer) {
// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 		time.Sleep(1 * time.Second)
// 	}

// 	fmt.Fprint(out, finalWord)
// }

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout , sleeper)
}
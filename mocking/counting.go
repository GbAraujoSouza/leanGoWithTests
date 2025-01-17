package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	CountDownStart = 3
	write          = "write"
	sleep          = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpyCountdownOperations struct {
	Calls []string
}

// implements Sleeper
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// implemets io.Writer
func (s *SpyCountdownOperations) Write(b []byte) (int, error) {
	s.Calls = append(s.Calls, write)

	var n int
	var err error
	return n, err
}

type DefaultSleeper struct{}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
  c.sleep(c.duration) 
}

type SpyTime struct {
  durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
  s.durationSlept = duration
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := CountDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(out, "Vai Teia!")
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, &sleeper)
}

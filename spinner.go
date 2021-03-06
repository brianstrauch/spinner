package spinner

import (
	"fmt"
	"time"
)

var frames = []string{"|", "/", "-", "\\"}

type Spinner struct {
	stop chan bool
	wait chan bool
}

func New() *Spinner {
	return &Spinner{
		stop: make(chan bool),
		wait: make(chan bool),
	}
}

func (s *Spinner) Start() {
	hideCursor()
	go s.run()
}

func (s *Spinner) Stop() {
	close(s.stop)
	<-s.wait
}

func (s *Spinner) run() {
	tick := time.Tick(time.Second / 3)
	for i := 0; i < len(frames); i = (i + 1) % len(frames) {
		select {
		case <-tick:
			clear()
			fmt.Print(frames[i])
		case <-s.stop:
			clear()
			showCursor()
			close(s.wait)
			return
		}
	}
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func clear() {
	fmt.Print("\r")
}

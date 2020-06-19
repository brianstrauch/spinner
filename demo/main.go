package main

import (
	"fmt"
	"time"

	"github.com/brianstrauch/spinner"
)

func main() {
	s := spinner.New()

	fmt.Println("Start")
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
	fmt.Println("Stop")
}

package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tickTime := range ticker.C {
		fmt.Println(tickTime)
	}
}

func TestTickerTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tickTime := range channel {
		fmt.Println(tickTime)
	}
}

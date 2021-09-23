package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())

	timeChannel := <-timer.C
	fmt.Println(timeChannel)
}

func TestTimeAfter(t *testing.T) {
	var channel = time.After(5 * time.Second)
	fmt.Println(time.Now())

	timeChannel := <-channel
	fmt.Println(timeChannel)
}

func TestTimeAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())

	group.Wait()
}

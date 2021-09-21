package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("jalan duluan?")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	fmt.Println("sending..")
	channel <- "Adiputra"
}

func OnlyOut(channel <-chan string) {
	fmt.Println("receiving..")
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go OnlyIn(channel)

	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 4)
	defer close(channel)

	fmt.Println("buffer size", cap(channel))
	fmt.Println("current buffer size", len(channel))

	channel <- "Adiputra"

	fmt.Println("current buffer size", len(channel))

	channel <- "Adiputra 2"

	fmt.Println("current buffer size", len(channel))

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// fmt.Println(<-channel)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data dengan range", data)
	}
}

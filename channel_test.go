package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	channel <- "Eko"

	data := <-channel

	fmt.Println(data)

	defer close(channel)
}

func TestCreateChannel2(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Mulai mengirim")
		channel <- "Paket.."
		fmt.Println("Blocking")
	}()

	fmt.Println("sebelum mengambil channel..")
	data := <-channel
	fmt.Println(data)
	fmt.Println("mengambil channel..")

	time.Sleep(5 * time.Second)

	defer close(channel)
}

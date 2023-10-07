package sesi2concurrency

import (
	"fmt"
	"runtime"
	"testing"
)

func process1(process chan string) {
	msg := "Process ke-1"
	process <- msg
}

func process2(process chan string) {
	msg := "Process ke-2"
	process <- msg
}

func process3(process chan string) {
	msg := "Process ke-3"
	process <- msg
}

func TestChannel(t *testing.T) {
	fmt.Printf("Running goroutine in %d cpu\n", runtime.NumCPU())

	process := make(chan string)

	go process1(process)
	go process2(process)
	go process3(process)

	fmt.Println("Value:", <-process)
	fmt.Println("Value:", <-process)
	fmt.Println("Value:", <-process)
}

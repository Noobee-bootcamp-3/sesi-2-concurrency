package sesi2concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func processX(ch chan string, kind string, done chan bool, total chan int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Running process X")

	fmt.Println("send data kind to ch")
	ch <- kind

	fmt.Println("get data from total and insert to totalData")
	totalData := <-total
	fmt.Println("Total data send:", totalData)
	totalData += 100

	done <- true
	fmt.Println("send data totalData to result channel")
	result <- totalData

	fmt.Println("Process X done...")
}

func processY(ch chan string, total chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Runnning process Y")

	fmt.Println("get data kind from ch channel")
	kind := <-ch

	switch kind {
	case "+":
		total <- 100 + 10
	case "-":
		total <- 100 - 10
	}
	fmt.Println("Process Y done...")
}

func TestChannelFunction(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan string)
	total := make(chan int)
	done := make(chan bool)
	result := make(chan int)

	go processX(ch, "-", done, total, result, &wg)
	go processY(ch, total, &wg)

	if <-done {
		fmt.Println("Total result recieve:", <-result)
	}

	wg.Wait()

	fmt.Println("Finish process..")
}

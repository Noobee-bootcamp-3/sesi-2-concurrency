package sesi2concurrency_test

import (
	"fmt"
	"sync"
	"testing"
)

func calculate(number ...int) (total int) {
	for _, num := range number {
		total += num
	}
	return
}

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}

	total := 0
	number := []int{1, 2, 3, 4, 5}

	wg.Add(1)

	go func(nums ...int) {
		total = calculate(nums...)
		wg.Done()
	}(number...)

	wg.Wait()

	fmt.Println("Total:", total)
}

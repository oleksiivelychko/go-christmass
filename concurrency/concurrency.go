package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

/*
A goroutine is a lightweight thread managed by the Go runtime.
A goroutine an abstraction over threads, a single OS thread can run many goroutines.
A goroutine takes about 2kb of stack space to initialize.
Goroutines run in the same address space, so access to shared memory must be synchronized.
Goroutine are not hardware dependent and does not have ID because go does not have Thread Local Storage.
*/
func goroutine(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
Goroutines have easy communication (with low latency) medium known as channel.
By default, sends and receives block until the other side is ready.
This allows goroutines to synchronize without explicit locks or condition variables.
*/
func channel(ch chan int, s []int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum // send sum to channel
}

/*
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.
*/
func selectChannel(tickCh chan string, quitCh chan int) {
	init := "init"
	for {
		select {
		case tickCh <- init:
			fmt.Println("tick")
		case <-quitCh:
			fmt.Println("quit")
			return // loop exit
		}
	}
}

func waitWorkers(wg *sync.WaitGroup) {
	for i := 1; i <= 3; i++ {
		wg.Add(1) // equals to count (2 times) of goroutines would be run

		go func(id int) {
			defer wg.Done() // alerting the WaitGroup when a goroutine completes

			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}
}

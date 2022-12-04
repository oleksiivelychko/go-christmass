package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	// determines how many threads that program will use simultaneously
	// swapping between threads is a relatively slow operation
	runtime.GOMAXPROCS(2)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("[func goroutine(s string)] -> %s:", r)
		}
	}()

	go goroutine("async")
	goroutine("sync")
}

/*
Non-buffered channels are blocking, meaning that when sending data into a channel, Go waits until the data is read from the channel before execution continues.
Buffered channels, meaning sending data into that channel won’t block until exceed the capacity.
*/
func TestChannel(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int, 2) // make buffered channel

	go channel(ch, s[:len(s)/2])
	go channel(ch, s[len(s)/2:])

	x, y := <-ch, <-ch

	close(ch) // optional
	_, ok := <-ch
	if ok == true {
		fmt.Println("channel has not been closed")
	}

	if x != -5 {
		t.Errorf("[func channel(ch chan int, s []int)] -> %d != -5", x)
	}

	if y != 17 {
		t.Errorf("[func channel(ch chan int, s []int)] -> %d != 17", y)
	}
}

func TestSelectChannel(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("[func selectChannel(tickCh chan string, quitCh chan int)] -> %s:", r)
		}
	}()

	tickCh := make(chan string)
	quitCh := make(chan int)

	go func() {
		fmt.Println(<-tickCh)
		fmt.Println(<-tickCh)
		quitCh <- 0
	}()

	selectChannel(tickCh, quitCh)
}

func TestMutex(t *testing.T) {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 3; i++ {
		go c.Inc("counter")
	}

	time.Sleep(time.Second)
	if c.Value("counter") != 3 {
		t.Errorf("[func (c *SafeCounter) Value(key string) int]: %d != 3", c.Value("counter"))
	}
}

func TestWaitGroups(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("[func waitWorkers(wg *sync.WaitGroup)] -> %s:", r)
		}
	}()

	// is used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup

	waitWorkers(&wg)

	// block the execution until the WaitGroup counter goes back to 0, all the workers notified they’re done.
	wg.Wait()
}

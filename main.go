package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create a new elevator system
	e := New(4, 8)

	// Create new observer for users
	observer := NewObserver()

	// Add channel for stopping execution
	stopChan := make(chan struct{})

	go func() {
		ticker := time.NewTicker(1000 * time.Millisecond)
		for {
			select {
			case <-stopChan:
				ticker.Stop()
				return
			case <-ticker.C:
				e.Step()
				observer.NotifyAll()
				for _, elevator := range e.Status() {
					fmt.Println(elevator)
				}
			}
		}
	}()

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Let the elevator system start up properly
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
			user := NewUser(e, rand.Intn(7))
			observer.Add(user)
			user.RequestElevator(rand.Intn(7))
			user.Block()
			observer.Remove(user)

			wg.Done()
		}()
	}

	wg.Wait()
	close(stopChan)
}

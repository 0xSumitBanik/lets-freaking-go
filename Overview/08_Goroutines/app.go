package main

import (
	"log"
	"math/rand"
	"time"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Goroutine is the Go way to create concurrent computations in Go programming.
	// Goroutines are also often called green threads.
	// Green threads are maintained and scheduled by the language runtime instead of the OS.
	// The cost of memory consumption and context switching, of a goroutine is much lesser than an OS thread.
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2) // Register two goroutines to wait for.
	go SayGreetings("hi!", 12)
	go SayGreetings("hello!", 10)
	time.Sleep(2 * time.Second)
	// When the main goroutine exits, the whole program also exits, even if there are still some other goroutines which have not exited yet.
	// How to let the main goroutine know when the two new goroutines have both finished their tasks?
	// We must use something called concurrency synchronization techniques.
	// here we will use another technique, the WaitGroup type in the sync standard package, to synchronize the executions between the two new goroutines and the main goroutine.
	wg.Wait() // blocks until the WaitGroup counter is zero.
}

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // sleep for 0 to 2.5 seconds
	}
	// Notify a task is finished.
	wg.Done() // ==  wg.Add(-1) subtracts one from the WaitGroup counter.
}

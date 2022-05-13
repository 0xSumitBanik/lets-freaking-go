package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	// Goroutine is the Go way to create concurrent computations in Go programming.
	// Goroutines are also often called green threads.
	// Green threads are maintained and scheduled by the language runtime instead of the OS.
	// The cost of memory consumption and context switching, of a goroutine is much lesser than an OS thread.
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	go SayGreetings("hi!", 12)
	go SayGreetings("hello!", 10)
	time.Sleep(2 * time.Second)
	// When the main goroutine exits, the whole program also exits, even if there are still some other goroutines which have not exited yet.
}

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // sleep for 0 to 2.5 seconds
	}
}

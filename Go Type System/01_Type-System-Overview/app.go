package main

// Basic types are covered under the Overview Section

/* Go supports the following composite types:
- pointer types - C pointer alike.
- struct types - C struct alike.
- function types - functions are first-class types in Go.
	> container types:
		- array types - fixed-length container types.
		- slice type - dynamic-length and dynamic-capacity container types.
		- map types - maps are associative arrays (or dictionaries).
		The standard Go compiler implements maps as hashtables.
		- channel types - channels are used to synchronize data among goroutines (the green threads in Go).
		- interface types - interfaces play a key role in reflection and polymorphism. */

var a *int        // a pointer type
var b [5]int      // an array type
var c []int       // a slice type
var d map[int]int // a map type

// a struct type
type info struct {
	name string
	age  int
}

// a function type
func abc(int) (bool, string)

// an interface type
type xyz interface {
	Method0(string) int
	Method1() (int, bool)
}

// some channel types
type def chan int
type ghi chan<- int

func main() {

}

package main

import "fmt"

/* Go doesn't support exception throwing and catching,
instead explicit error handling is preferred to use in Go programming.
In fact, Go supports an exception throw/catch alike mechanism.
The mechanism is called panic/recover.
*/

func main() {
	defer func() {
		fmt.Println("exit normally.")
	}()
	fmt.Println("Hello")
	defer func() {
		// The value returned by a recover function call is the value a panic function call consumed.
		v := recover()
		fmt.Println("recovered:", v)
	}()
	panic("bye!")
	fmt.Println("This code is unreachable, as panic() is called before this statement.")
}

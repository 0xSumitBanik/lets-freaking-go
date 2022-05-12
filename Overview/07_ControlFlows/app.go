package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Note: basic control flow code blocks: if-else, for, switch-case
// Control Flow Code Blocks related to certain type: for-range, select-case, type-switch

func main() {

	rand.Seed(time.Now().UnixNano())

	// Example of if-else code block
	if x := rand.Int(); x&1 == 1 {
		fmt.Printf("%v is Odd number\n", x)
	} else { // else has to be in the same line with the closing brace of if{} ✅
		fmt.Printf("%v is Even number\n", x)
	}

	// Example: for loop control flow
	// Print 10 numbers
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v, ", i)
	}

	// Alternate way to print 10 numbers
	var x int = 1
	for ; x <= 10; x++ {
		fmt.Printf("%v, ", x)
	}

	// Using for loop as a while loop
	j:=1
	for j<=10{
		fmt.Printf("%v, ", j)
		j++
	}
}

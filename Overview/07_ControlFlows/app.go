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
	j := 1
	for j <= 10 {
		fmt.Printf("%v, ", j)
		j++
	}
	// The rand.Intn function returns a non-negative int random value which is smaller than the specified argument.
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n%9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")

		// Different from many other languages,
		// in Go, the execution will automatically
		// jumps out of the switch-case block at
		// the end of each branch block.
		// No "break" statement is needed here.
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		// here, this "break" statement is redundant.
		break
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	// case 6, 7, 8:
		// The above case line might fail to compile,
		// for 6 is duplicate with the 6 in the last
		// case. The behavior is compiler dependent.
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}

	// To allow the code execution to move into next case without breaking we can use fallthrough keyword
	// fallthrough has to be the final statement for any case branch
	
}

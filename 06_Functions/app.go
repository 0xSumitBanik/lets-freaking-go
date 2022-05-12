package main

import "fmt"

// There can be multiple functions named as init declared in a package, even in a source code file.
// The functions named as init must have not any input parameters and return results.
func init() {
	fmt.Println("This will be printed first")
}

func add(x int, y int) int {
	return x + y
}

// If the type portions of some successive parameters in a parameter declaration list are identical,
// then these parameters could share the same type portion in the parameter declaration list.
// 👆 Same goes for the return variable type.

func addAndSub(x, y int) (s, d int) {
	return x + y, x - y
	// Named return:
	// s = x + y
	// d = x - y
	// return
}

func main() {
	var res int = add(1, 2)
	addn, sub := addAndSub(1, 26)
	println(res)
	println(addn, sub)

	// This anonymous function has no parameters
	// but has two results.
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // Call it. No arguments are needed.

	// The following anonymous function have no results.

	func(a, b int) {
		// The following line prints: a*a + b*b = 25
		println("a*a + b*b =", a*a+b*b)
	}(x, y) // pass argument x and y to parameter a and b.

	func(x int) {
		// The parameter x shadows the outer x.
		// The following line prints: x*x + y*y = 32
		println("x*x + y*y =", x*x+y*y)
	}(y) // pass argument y to parameter x.

	func() {
		// The following line prints: x*x + y*y = 25
		println("x*x + y*y =", x*x+y*y)
	}() // no arguments are needed.

}

// An init function in an importing package will be invoked after all the init functions declared in the dependency packages of the importing package for sure. 
func init() {
	fmt.Println("This will be printed second")
}

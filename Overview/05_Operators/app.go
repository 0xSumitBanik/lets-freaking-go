package main

func main() {
	// Arithmetic operators
	// +, -, *, /, %, ++, --
	println("Arithmetic operators")
	println(10 + 10)
	println(10 - 10)
	println(10 * 10)
	println(10 / 10)
	println(10 % 10)

	// println(20.5%1) // This will give an error as mod is not supported in floating numbers.

	// Binary Arithmetic Operators
	//  &, |, ^, <<, >>, &^
	println("Binary Arithmetic Operators")
	println(10 & 10)
	println(10 | 10)
	println(10 ^ 10)
	println(10 << 10)
	println(10 >> 10)
	println(10 &^ 10) // = 10 & (^10)

	// Unary Operators
	// +, -,  ^
	println("Unary Operators")
	println(+10)
	println(-10)
	var x uint = 10
	println(^x)

	// Logical Operators
	print(true && false)
	print(false || true)
	print(!true)

	// Comparison Operators
	// Go supports six comparison binary operators:
	// ==, !=, <, >, <=, and >=

	/* Operator Precedence

	*   /   %   <<  >>  &   &^
	+   -   |   ^
	==  !=  <   <=  >   >=
	&&
	||
	*/

}

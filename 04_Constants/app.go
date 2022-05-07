package main

func main() {

	const i = 5
	// i=10 👈 This will give an error
	// Note: If we don't use the constants anywhere in the program, it won't throw error.
	// But if we use declared unused variables in the program, it will throw error.
	// var j int = 100
	// 👆If we don't use the variable, it'll give error: .\app.go:10:6: j declared but not used
	//boolean: false and true are named constants

	// Explicit Conversions
	// Go also supports value conversions. We can use the form Type(v) to convert a value v

	// Example:
	/*
		float32(123)
		uint(1.0)
		int8(-123)
		int16(6+0i)
		complex128(789)

		string(65)          // "A"
		string('A')         // "A"
		string('\u68ee')    // "森"
		string(-1)          // "\uFFFD"
		string(0xFFFD)      // "\uFFFD"
		string(0x2FFFFFFFF) // "\uFFFD"

	*/

	// Declare two individual constants.
	// non-ASCII letters can be used in identifiers.
	const π = 3.1416
	const Pi = π // <=> const Pi = 3.1416

	// Constant Grouping.
	const (
		no = false

		a  = 1
		b  // this will hold the value of the constant declared above
		ch = "Hello"
	)
	println(no, a, b, ch)

	// Typed constants

	const (
		ii, jj int = 123, 234
		ab, _      // ab = 123, _ = 234
		// jj int = 345
	)

	// iota in constant declarations
	// iota is a predefined identifier used to generate unique values for constants.
	// It is declared in the same scope as the constant.
	// The value of iota is 0 at the start of each new constant declaration.
	// The value of iota is incremented for each new constant.
	const (
		kk = 100       // here the iota=0
		ll = iota      // ll=1
		mm = iota + 10 // mm = 2+10 = 12
	)

	// The value denoted by an untyped constant can overflow its default type
	// Example
	// 3 untyped named constants. Their bound
	// values all overflow their respective
	// default types. This is allowed.
	const n = 1 << 64          // overflows int
	const r = 'a' + 0x7FFFFFFF // overflows rune
	const x = 2e+308           // overflows float64
}

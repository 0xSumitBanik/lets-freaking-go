package main

func main() {
	// The var statement declares a list of variables
	// The type is mentioned at the last.
	// Note: If we have multiple variables declared, it's  good to declare in factored form.

	var (
		flag bool // boolean data type
		i int		// integer data type
		j uint	// unsigned integer data type
		k float32	// float data type
		s string	// string data type
		b byte // alias for uint8
		r rune // alias for int32
	)

	println(flag, i, j, k, s, b, r) // This will print falsy (default) values for all the variables.
	// Let's set values to the above variables
	flag = true
	i = -10
	j = 20
	k = 3.14
	s = "Hello"
	b = 'a' // This will hold the ascii value of 'a'
	r = 'b' // This will hold the ascii value of 'b'
	println(flag, i, j, k, s, b, r) // This will print true, -10, 20, 3.14, Hello, 97, 98

	x:= 10 // This is a short hand operator for var x int = 10 (we don't need to specify the data type)
	y:= "HI" // This is a short hand operator for var y string = "HI"

	println(x,y)

	// Integer Value Literals
	// The following are the different types of integer value literals:
	// 1. Decimal
  var dd int = 28
	// 2. Hexadecimal
	var hh int = 0x1f
	// 3. Octal
	var oo int = 027 // prefix it with 0
	// 4. Binary
	var bb int = 0b1010

	println(dd, hh, oo, bb)

	// Floating Point Value Literals
	// The following are the different types of floating point value literals:
	// 1. Decimal
	var dd1 float32 = 28.5
	// 2. Hexadecimal
	var hh1 float32 = 0x1f.p5
	// 3. Octal
	var oo1 float32 = 027.5 // prefix it with 0
	println(dd1, hh1, oo1)
	
	// String Value literals
	// The following are the different types of string value literals:
	// 1. Double quoted string
	var s1 string = "Hello\nWorld"
	// 2. Raw Form String
	var s2 string = `Hello
	World`

	println(s1, s2)
}
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
}
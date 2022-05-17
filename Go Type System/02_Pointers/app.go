package main

var x *int  // An unnamed pointer type whose base type is int.
var y **int // An unnamed pointer type whose base type is *int.

// Ptr is a named pointer type whose base type is int.
type Ptr *int
var tx Ptr
// PP is a named pointer type whose base type is Ptr.
type PP *Ptr

// All of the above pointers store the value as <nil> initially.

// The built-in new function can be used to allocate memory for a value of any type.
var xx *int = new(int) // xx is a pointer to an int which holds a memory value.

var jj *int

// We cannot get the value of *jj (as it is a nil pointer), but we can get the type of jj.

// return the address of a local variable is absolutely safe in Go.
func newInt() *int {
	a := 3
	return &a
}

//  We can compare pointers that are of equivalent types
var res bool = (x == tx) // true

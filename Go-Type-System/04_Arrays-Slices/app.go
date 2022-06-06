package main

import (
	"fmt"
)

// Arrays,Slices and Maps are first-class citizen container types in Go.
// Each element in a container has an associated key.
// For arrays and slices, it's an int value.
// For maps, the key is of any comparable type.

const Size = 32

type Person struct {
	name string
	age  int
}

func main() {

	/* Array types */

	var array [Size]int
	declaredArray := [5]int{1, 1: 2, 3, 4, 5} //we can also mention the index as key (optional)
	// Element type is a struct type: Person
	var personArray [4]Person
	fmt.Printf("Type: %T %T\n", array, personArray)
	fmt.Println(personArray, "\n\n Declared Array ->", declaredArray)
	fmt.Printf("\n The length of array is %v`\n", len(array))

	/* Slice types */

	var slice []int
	var personSlice []Person
	fmt.Printf("Type: %T %T\n", slice, personSlice)
	declaredSlice := []Person{{name: "name1", age: 1}, {name: "name2", age: 2}}
	fmt.Println(personSlice, "\n\n Declared Slice ->", declaredSlice)
	fmt.Printf("\n The length of slice is %v\n", len(declaredSlice))
	fmt.Printf("\n The capactiy of declaredslice is %v\n", cap(declaredSlice))

	// A slice is an abstraction that uses an array under the covers.
	newSlice := make([]int, 3) // newSlice = [0,0,0]

	for i := 0; i < 5; i++ {

		// It will double the capacity for the next 3 elements and will create the array with the new capacity and memory address.
		newSlice = append(newSlice, i)
		fmt.Printf("\n Capacity %v, Length %v, Slice Address: %p, %v", cap(newSlice),
			len(newSlice),
			newSlice,
			newSlice)
		/* Output:
		Capacity 6, Length 4, Slice Address: 0xc00000c690, [0 0 0 0]
		Capacity 6, Length 5, Slice Address: 0xc00000c690, [0 0 0 0 1]
		Capacity 6, Length 6, Slice Address: 0xc00000c690, [0 0 0 0 1 2]
		Capacity 12, Length 7, Slice Address: 0xc00001c240, [0 0 0 0 1 2 3]
		Capacity 12, len 8, Slice Address: 0xc00001c240, [0 0 0 0 1 2 3 4]
		*/
	}

	copiedSlice := make([]int, len(newSlice))
	copy(copiedSlice, newSlice) // copy(destination,source)
	fmt.Printf("\n Copied Slice %v , Address %p -> ", copiedSlice, copiedSlice)

	rangeSlice := copiedSlice[4:7] // rangeSlice = [4,5,6]
	rangeSlice[0] = 100
	fmt.Printf("\n Range Slice %v , Address %p -> ", rangeSlice, rangeSlice)

	// the ...s mean we want to let compilers deduce the lengths for the corresponding array values.
	variableArray := [...]int{}
	fmt.Printf("\n Variable Array %T \n", variableArray) // [0]int

	ps0 := &[]string{"Go", "C"} // Pointer to the string slice
	fmt.Printf("\n Pointer to a string %v", *ps0)

	// Ref: https://go101.org/article/container.html

	fmt.Println("\n Looping through Array/slice using range function")
	// Looping through Array/slice using range function
	for j := range(newSlice){
		fmt.Println(newSlice[j])
	}
}

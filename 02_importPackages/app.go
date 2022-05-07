package main

// To use custom packages we use the import statement.
// Ref: https://go.dev/ref/spec#Import_declarations

import (
	"fmt"
	"math"
)

// 👆 It is good style to use the factored import statement.

func main(){
	// The form aImportName.AnExportedIdentifier is called a qualified identifier
	fmt.Println("Square root of 25 is",math.Sqrt(25))
}

// Note: The Sqrt() function is capitalized which means it is an "exported function" (from math package)
//       If we want to make a function available outside the package, we capitalize the function name.
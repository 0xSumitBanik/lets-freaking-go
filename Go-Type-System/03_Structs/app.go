package main

import "fmt"

// structure field names are also called as member variables.
type Book struct {
	title  string
	author string
	pages  int
}

// Consecutive fields with the same type can be declared together.
type BookNew struct {
	title,author  string
	pages  int
}

// Size(struct) = sum(all field types)+padding bytes

// Topic: Struct Field Tags
// A tag may be bound to a struct field when the field is declared.
type BookExport struct {
	Title  string `json:"title" customTag:"customValue"`
	author string
	pages  int
}

// 👆 In the above example, the field tags can help the functions 
// in the encoding/json standard package to determine the field names in JSON texts
// It will only encode and decode the exported struct fields (capitalized field names)

func main(){
	be := BookExport{}
	fmt.Println(be) // {  0} (two empty strings and 0)
	// Initialize a struct value using selectors
	be.Title="Demo"
	fmt.Println(be) // {Demo  0} (Demo, empty string and 0)
}
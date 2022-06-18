package main

import (
	"fmt"
)

func main(){
	// var x map[string]int
	// x["name"]=12 // panic: assignment to entry in nil map
	y := make(map[string]int)
	y["zero"]=0
	y["one"]=1
	y["two"]=2
	y["three"]=3
	fmt.Println(y) // map[one:1 two:2 three:3 zero:0] (printed keeping the key in sorted order)

	value,isPresent := y["zero"]
	fmt.Println(value,isPresent) // 0 true
	value,isPresent = y["four"]
	fmt.Println(value,isPresent) // 0 false
	delete(y,"two")
	value,isPresent = y["two"]
	fmt.Println(value,isPresent) // 0 false
	// looping through map
	for k,v:=range y{
		fmt.Println(k,v) // one 1, two 2, three 3, zero 0
	}

	newMap := make(map[string]map[string]int)
	newMap["one"]=make(map[string]int)
	fmt.Println(newMap)
}
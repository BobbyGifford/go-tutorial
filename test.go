package main

import "fmt"

var decksize int

var test = 70

func testFun(testVar int) (int, string) {
	return (testVar * 2), "Go can have two return values"
}

func main() {
	// innerTest := "New var"
	// fmt.Println(test)
	// fmt.Println(innerTest)

	testVar1 := "this is valid"

	fmt.Printf("testVar is of type %T \n", testVar1)

	fmt.Println(testFun(10))
}

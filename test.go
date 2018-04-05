package main

import "fmt"

var decksize int

var test = 70


func testFun(testVar int) (int, string) {
	return (testVar * 2), "go can have two return values"
}

func main() {
	innerTest := "New var"
	fmt.Println(test)
	fmt.Println(innerTest)
}

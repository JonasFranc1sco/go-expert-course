package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = "Hello"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("Type of var is %T and value is %v\n", t, t)
}

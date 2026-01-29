package main

import "fmt"

func main() {
	var myVar interface{} = "Jonas"
	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("Result value is %v and result of ok is %v\n", res, ok)
	res2 := myVar.(int) // Result: panic: interface conversion: interface {} is string, not int
	fmt.Printf("Res2 value is %v", res2)
}

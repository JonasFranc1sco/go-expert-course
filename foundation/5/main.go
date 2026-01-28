package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 9
	d string  = "Jonas"
	e float64 = 1.7
	i ID      = 1
)

func main() {
	var MyArray [3]int // 3 fixed positions
	MyArray[0] = 10
	MyArray[1] = 20
	MyArray[2] = 30

	for i, v := range MyArray {
		fmt.Printf("The index value %d is %d\n", i, v)
	}
}

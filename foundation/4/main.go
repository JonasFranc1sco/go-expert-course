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
	fmt.Printf("O tipo de E é %T", i) // T -> Type of
}

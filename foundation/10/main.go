package main

import "fmt"

func main() {

	// Closures: A function can have multiple functions
	total := func() int {
		return sum(1, 131, 123, 123, 46, 564, 654, 65, 657, 345, 234, 123) * 2
	}()
	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

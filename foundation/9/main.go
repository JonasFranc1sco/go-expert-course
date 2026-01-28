package main

import "fmt"

func main() {
	fmt.Println(sum(1, 131, 123, 123, 46, 564, 654, 65, 657, 345, 234, 123))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

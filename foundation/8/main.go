package main

import (
	"errors"
	"fmt"
)

func main() {

	// First code with error handling
	value, err := sum(10, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

// Need parameters a and b, return an int and error
func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("The sum is greater than 50.")
	}
	return a + b, nil
}

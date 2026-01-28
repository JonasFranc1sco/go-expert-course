package main

import "fmt"

func main() {
	salaries := map[string]int{"Wesley": 1000, "Jonas": 2000, "Marcia": 5000}
	fmt.Println(salaries["Jonas"])
	delete(salaries, "Wesley")
	fmt.Println(salaries)
	salaries["Eduardo"] = 4000
	fmt.Println(salaries["Eduardo"])

	for name, salary := range salaries {
		fmt.Printf("Salary of %s is %d\n", name, salary)
	}

	for _, salary := range salaries { // _ = blank identifier
		fmt.Printf("Salary: %d\n", salary)
	}
	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Jonas"] = 1000
}

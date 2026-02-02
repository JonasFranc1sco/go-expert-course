package main

import (
	"21/mathematic"
	"fmt"

	"github.com/google/uuid"
)

// Todas as coisas com letra minúscula (Var, Struct, Func) = Internas
// Todas as coisas com letra maiúscula (Var, Struct, Func) = Externas

func main() {
	s := mathematic.Sum(10, 20)
	car := mathematic.Car{Brand: "Ford"}
	fmt.Printf("Result: %v\n", s)
	fmt.Println(mathematic.A)
	fmt.Println(car)
	fmt.Println(uuid.New())
}

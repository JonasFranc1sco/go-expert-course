package main

type MyNumber int

type Number interface {
	~int | float64 // ~ = Consider use other type ints
}

func SumInt(m map[string]int) int {
	var sum int
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumFloat(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

// Transforming Sum function in generics
// T = Int or float
func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Jonas": 1000, "Marcia": 2000, "João": 3000}
	m2 := map[string]float64{"Jonas": 10.00, "Marcia": 200.000, "João": 3.000}
	m3 := map[string]MyNumber{"Jonas": 1000, "Marcia": 2000, "João": 3000}

	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
	println(Compare(10, 1.0))
}

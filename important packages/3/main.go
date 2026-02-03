package main

import (
	"fmt"
	// "io"
	// "net/http"
)

// func main() {
// 	req, err := http.Get("https://www.google.com")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer req.Body.Close() // Just run after all ends
// 	res, err := io.ReadAll(req.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(res))
// }

func main() {
	defer fmt.Println("First line.")
	fmt.Println("Second line.")
	fmt.Println("Third line.")

}

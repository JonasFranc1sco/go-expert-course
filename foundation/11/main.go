package main

import "fmt"

type Client struct {
	Name   string
	Year   int
	Active bool
}

func main() {
	jonas := Client{
		Name:   "Jonas",
		Year:   20,
		Active: true,
	}

	fmt.Printf("Name: %s, Year: %d, is Active: %t\n", jonas.Name, jonas.Year, jonas.Active)
}

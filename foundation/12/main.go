package main

import "fmt"

type Localization struct {
	Number     int
	Street     string
	City       string
	State      string
	Complement string
}

type Client struct {
	Name    string
	Year    int
	Active  bool
	Address Localization
}

//type Client struct {
//	Name    string
//	Year    int
//	Active  bool
//	Localization
//}

func main() {
	jonas := Client{
		Name:   "Jonas",
		Year:   20,
		Active: true,
	}

	jonas.Address.City = "Porto Velho"

	fmt.Printf("Name: %s, Year: %d, is Active: %t\n", jonas.Name, jonas.Year, jonas.Active)
}

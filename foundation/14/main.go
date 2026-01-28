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

type User interface {
	Deactivate()
}

type Enterprise struct {
	Name string
}

func (e Enterprise) Deactivate() {
}

// Method for Struct
func (c *Client) Deactivate() {
	c.Active = false
	fmt.Printf("Client %s was deactivated\n", c.Name)
}

func Deactivation(user User) {
	user.Deactivate()
}

func main() {
	jonas := Client{
		Name:   "Jonas",
		Year:   20,
		Active: true,
	}
	myEnterprise := Enterprise{}

	jonas.Deactivate()
	Deactivation(myEnterprise)
	fmt.Printf("Name: %s, Year: %d, is Active: %t\n", jonas.Name, jonas.Year, jonas.Active)
}

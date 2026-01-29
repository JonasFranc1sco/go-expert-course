package main

import "fmt"

type Client struct {
	balance int
}

func NewClient() *Client {
	return &Client{balance: 0}
}

// Without pointer
//
//	|
//	V
func (c Client) simulate(value int) int { // Just a copy, not a memory address value
	c.balance += value
	fmt.Printf("Simulate balance value: %d\n", c.balance)
	return c.balance
}

// With pointer
//
//	|
//	V
func (c *Client) deposit(value int) int { // Pointer literal balance memory address, not a copy
	c.balance += value
	fmt.Printf("New balance: %d\n", c.balance)
	return c.balance
}

func main() {
	jonas := Client{
		balance: 100,
	}
	fmt.Printf("Initial balance: %d\n", jonas.balance)
	jonas.deposit(200)
	jonas.simulate(1000)
	fmt.Printf("Your balance value: %d\n", jonas.balance)
}

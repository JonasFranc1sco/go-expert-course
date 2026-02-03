package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	lenght, err := f.Write([]byte("Writing data..."))
	//lenght, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File successful created! Lenght: %d bytes.\n", lenght)
	f.Close()

	// Read

	file, err := os.ReadFile(("file.txt"))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Read step by step file
	file2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}

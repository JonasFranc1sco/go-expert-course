package main

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 9
	d string  = "Jonas"
	e float64 = 1.7
	i ID      = 1
)

func main() {
	n := 2
	n = 10
	println(n)
	id()
}

func id() {
	println(i)
}

func x() {

}

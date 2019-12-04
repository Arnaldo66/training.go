package data

import "fmt"

var Name = "Bob"
var Age = 20
var password = "secretpassword"

func Greeting(name string)  {
	fmt.Printf("%s va fan culo\n", Name)
}

func Avg(x, y float64) float64  {
	return (x + y) / 2
}

func SumNamed(x, y, z int) (sum int)  {
	sum = x + y + z
	return
}

func ManyReturn() (int, string)  {
	return 4, "hello"
}

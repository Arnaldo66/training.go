package main

import "fmt"

type User struct {
	Name string
}

func main() {
	myMap := map[string]*User{
		"HR": {"boby"},
		"CTO": {"jean"},
	}

	fmt.Println(myMap["ii"])
	fmt.Println(myMap["CTO"])

	i := myMap["CTO"]
	i.Name = "Arnol"

	fmt.Println(myMap["CTO"])
	myMap["TA"] = &User{"jjj"}
	fmt.Println(myMap["TA"])

}

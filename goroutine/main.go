package main

import (
	"fmt"

)

func write(c chan string)  {
	names := []string{"Arnaud", "Jean", "Malouda"}
	for _, name := range names{
		c <- name
		fmt.Printf("Wrote %v to channel (len: %v)\n", name, len(c))
	}
	close(c)
}

func main()  {
	c := make(chan string, 3)

	fmt.Printf("Channel data: cap=%d, len=%v\n", cap(c), len(c))

	go write(c)
	//time.Sleep(2 * time.Second)

	for v := range c {
		fmt.Println(v)
	}
}
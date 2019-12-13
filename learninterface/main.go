package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Println(err)
	}


	defer resp.Body.Close()

	body, _  := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
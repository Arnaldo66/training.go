package main

import (
	"fmt"
	"os"
	cert2 "training.go/gencert/cert"
	"training.go/gencert/pdf"
)

func main()  {
	cert, err := cert2.New("Golang programming", "Boby", "2018-01-02")
	if err != nil {
		fmt.Printf("Error during certificate")
		os.Exit(1)
	}

	var saver cert2.Saver
	saver, err = pdf.New("output")
	if err != nil {
		fmt.Printf("Error during generating pdf")
		os.Exit(1)
	}
	saver.Save(*cert)

}

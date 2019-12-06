package main

import(
    "fmt"
    "os"
    "training.go/dictionary/dictionary"
)

func main()  {
    d, err := dictonary.New("./badger")
    handleErr(err)
    defer d.Close()
}

func handleErr(err error)  {
    if err != nil {
        fmt.Printf("Dictionary error: %v", err)
        os.Exit(1)
    }
}

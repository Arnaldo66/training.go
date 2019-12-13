package main

import(
    "fmt"
    flag "github.com/spf13/pflag"
    "os"
    "training.go/dictionary/dictionary"
)

func main()  {
    action := flag.String("action", "list", "Action to perform on the dictionary")

    d, err := dictonary.New("./badger")
    handleErr(err)
    defer d.Close()

    flag.Parse()
    switch *action {
    case "list" :
        actionList(d)
    case "add":
        actionAdd(d, flag.Args())
    case "remove":
        actionRemove(d, flag.Args())
    default:
        fmt.Println("Unknow action")
    }
}
func actionRemove(d *dictonary.Dictionary, args []string)  {
    word := args[0]
    err := d.Remove(word)
    handleErr(err)
    fmt.Println("Le mot a bien été supprimé")

}

func actionAdd(d *dictonary.Dictionary, args []string)  {
    word := args[0]
    definition := args[1]
    err := d.Add(word, definition)
    handleErr(err)
    fmt.Println("Le mot a bien été enregistré")
}

func actionList(d *dictonary.Dictionary)  {
    words, entries, err := d.List()
    handleErr(err)
    fmt.Println("Dictionary content")
    for _, word := range words {
        fmt.Println(entries[word])
    }
}

func handleErr(err error)  {
    if err != nil {
        fmt.Printf("Dictionary error: %v", err)
        os.Exit(1)
    }
}



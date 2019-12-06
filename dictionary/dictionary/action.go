package dictonary

import (
    "time"
    "bytes"
    "encoding/gob"
    "github.com/dgraph-io/badger"
)

func (d *Dictionary) Add(word, definition string) error  {
    entry := Entry{
        Word: word,
        Definition: definition,
        CreatedAt: time.Now(),
    }
    var buffer bytes.Buffer
    enc := gob.NewEncoder(&buffer)
    err := enc.Encode(entry)
    if err != nil {
        return nil
    }

    return d.db.Update(func (txn *badger.Txn) error  {
        return txn.Set([]byte(word), buffer.Bytes())
    })
}

# NANODB

A very tiny key-value store that can handle any number of inserts at the same time. 

# Usage

```go

db,err  := nanodb.Open("store.db")
if err != nil{
    log.Fatal(err)
}

// You will have to handle type conversions, the only input is string
key := "This is a key"
value := "this is a value"

err = db.Put(key, value)
if err != nil{
    log.Fatal(err)
}

val := db.Get(key)
fmt.Printf("%s", val)

err = db.Delete(key)
if err != nil{
    log.Fatal(err)
}
```

# Future Work

As of now the DB just stores it in file directly, this leads to a o(n) search for get and delete. Inserting the key value pairs in a sorted (by keys) will make it o(log n). Will do so if I get the time and interest. 

# License

MIT
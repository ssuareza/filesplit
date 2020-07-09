# About
This is a Go library to split files into pieces.

# Installing
```sh
$ go get github.com/ssuareza/filesplit
```

# Example
```go
import (
    "github.com/ssuareza/filesplit"
)

func main() {
    // split file
	chunks, err := filesplit.Split("/tmp/file.dat")
	if err != nil {
		log.Fatal(err)
    }
    
    // list chunks
	for _, chunk := range chunks {
		fmt.Println(chunk.Name)
    }
    
    // save it to disk
    filesplit.Save(chunks, "/tmp/"); err != nil {
		log.fatal(err)
	}
}
```
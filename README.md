# go-workshop

## part 1

### install

- Download and install Go https://go.dev/

- Install VS Code extension for Go

### print hello

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

`go run main.go`

### hello function

```go
func main() {
	hello()
}

func hello() {
	fmt.Println("Hello world")
}
```

### go mod init

`go mod init github.com/jensravn/go-workshop`

`go run .`

### hello server

```go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := []byte("Hello world")
		w.Write(b)
	})
	http.ListenAndServe(":8080", nil)
}
```

### get json

### get data type thing

### post

### db file

## part 2

### test endpoint

### router

### server

### dbFile

### thing interface

### test thing

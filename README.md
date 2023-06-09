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

```go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		j := map[string]any{"message": "Hello world"}
		s, err := json.Marshal(&j)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		b := []byte(s)
		w.Write(b)
	})
	http.ListenAndServe(":8080", nil)
}
```

### get type thing

```go
type thing struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := thing{
			Message: "Hello world",
		}
		s, err := json.Marshal(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		b := []byte(s)
		w.Write(b)
	})
	http.ListenAndServe(":8080", nil)
}
```

### put

```go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t := thing{
				Message: "Hello world",
			}
			s, err := json.Marshal(&t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			b := []byte(s)
			w.Write(b)
			return
		}
		if r.Method == http.MethodPut {
			var t thing
			err := json.NewDecoder(r.Body).Decode(&t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Printf("got thing: %#v", t)
			return
		}
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	})
	http.ListenAndServe(":8080", nil)
}
```

### db file

## part 2

### test endpoint

### router

### server

### dbFile

### thing interface

### test thing

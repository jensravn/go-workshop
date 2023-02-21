# Go workshop

## install

- Download and install Go https://go.dev/

- Install VS Code extension for Go

## print hello

```go
func main() {
	fmt.Println("hello")
}
```

run hello program: `go run main.go`

## hello function

```go
func hello() {
	fmt.Println("hello")
}
```

## go mod init

`go mod init github.com/USERNAME/go-workshop`

## hello server

```go
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
```

## json

```go
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	data["msg"] = "hello"
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not marshal data to json, err=%v", err)
	}
	w.Write(jsonData)
}
```

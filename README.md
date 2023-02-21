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

## data type thing

```go
func main() {
	http.HandleFunc("/thing", handleThing)
	http.ListenAndServe(":8080", nil)
}

type Thing struct {
	Msg string `json:"msg"`
}

func handleThing(w http.ResponseWriter, r *http.Request) {
	thing := Thing{
		Msg: "hello",
	}
	jsonThing, err := json.Marshal(thing)
	if err != nil {
		log.Printf("Could not marshal thing to json, err=%v", err)
	}
	w.Write(jsonThing)
}
```

## post

```go
func main() {
	http.HandleFunc("/thing", handleThing)
	http.ListenAndServe(":8080", nil)
}

type Thing struct {
	Msg string `json:"msg"`
}

func handleThing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getThing(w, r)
		return
	}
	if r.Method == "POST" {
		postThing(w, r)
		return
	}
	log.Printf("Error: /thing method '%s' not allowed", r.Method)
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func getThing(w http.ResponseWriter, r *http.Request) {
	thing := Thing{
		Msg: "hello",
	}
	jsonThing, err := json.Marshal(thing)
	if err != nil {
		log.Printf("Could not marshal data to json, err=%v", err)
	}
	w.Write(jsonThing)
}

func postThing(w http.ResponseWriter, r *http.Request) {
	var t Thing
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("r.Body could not be decoded into Thing struct, err=%v", err)
	}
	log.Printf("%#v", t)
}
```

## db file

```go
func main() {
	http.HandleFunc("/thing", handleThing)
	http.ListenAndServe(":8080", nil)
}

type Thing struct {
	Msg string `json:"msg"`
}

func handleThing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getThing(w, r)
		return
	}
	if r.Method == "POST" {
		postThing(w, r)
		return
	}
	log.Printf("Error: /thing method '%s' not allowed", r.Method)
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func getThing(w http.ResponseWriter, r *http.Request) {
	thing := dbGet()
	jsonThing, err := json.Marshal(thing)
	if err != nil {
		log.Printf("Could not marshal data to json, err=%v", err)
	}
	w.Write(jsonThing)
}

func postThing(w http.ResponseWriter, r *http.Request) {
	var t Thing
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("r.Body could not be decoded into Thing struct, err=%v", err)
	}
	dbPut(t)
}

func dbPut(t Thing) {
	bytes, err := json.Marshal(t)
	if err != nil {
		log.Printf("marshal err=%v", err)
	}
	os.WriteFile("thing.json", bytes, 0644)
}

func dbGet() *Thing {
	bytes, err := os.ReadFile("thing.json")
	if err != nil {
		log.Printf("readfile err=%v", err)
	}
	var t Thing
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		log.Printf("unmarshal err=%v", err)
	}
	return &t
}
```

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

```go
const thingTXT = "thing.txt"

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	var t thing
	if r.Method == http.MethodGet {
		b, err := os.ReadFile(thingTXT)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
		return
	}
	if r.Method == http.MethodPut {
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s, err := json.Marshal(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b := []byte(s)
		os.WriteFile(thingTXT, b, 0644)
		return
	}
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
```

## part 3

### test endpoint

```go
func Test_handleIndex(t *testing.T) {

	// given
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", strings.NewReader(``))

	// when
	handleIndex(w, r)

	// then
	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("status code got=%d, want=%d", res.StatusCode, 200)
	}
	got := string(body)
	want := `{"message":"test message"}`
	if got != want {
		t.Fatalf("body got=%s, want %s", got, want)
	}
}
```

### router

```go
func main() {
	r := chi.NewRouter()
	r.Get("/", handleGet)
	r.Put("/", handlePut)
	http.ListenAndServe(":8080", r)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile(thingTXT)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	var t thing
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s, err := json.Marshal(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := []byte(s)
	os.WriteFile(thingTXT, b, 0644)
}
```

### server

```go

func main() {
	s := server{}
	s.routes()
	http.ListenAndServe(":8080", s.r)
}

type server struct {
	r *chi.Mux
}

func (s *server) routes() {
	r := chi.NewRouter()
	r.Get("/thing", s.handleGet)
	r.Put("/thing", s.handlePut)
	s.r = r
}

func (s *server) handleGet(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile(thingTXT)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (s *server) handlePut(w http.ResponseWriter, r *http.Request) {
	var t thing
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	os.WriteFile(thingTXT, b, 0644)
}
```

### dbFile

```go
type dbFile struct{}

func (db *dbFile) Get() (*thing, error) {
	b, err := os.ReadFile(thingTXT)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	var t thing
	err = json.NewDecoder(bytes.NewReader(b)).Decode(&t)
	if err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}
	return &t, nil
}

func (db *dbFile) Put(t *thing) error {
	b, err := json.Marshal(&t)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}
	os.WriteFile(thingTXT, b, 0644)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}
```

### db interface

```go
type db interface {
	Get() (*thing, error)
	Put(*thing) error
}
```

### test thing

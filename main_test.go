package main

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_handleIndex(t *testing.T) {

	// given
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", strings.NewReader(``))

	// when
	handleGet(w, r)

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

package main

import (
	"fmt"
	"io"
	"net/http"
)

func MethodChecks(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "You made a %s request.", r.Method)

}

func echoText(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", 405)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if data == nil {
		http.Error(w, "empty file", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, string(data))
}

func HeadersDetective(w http.ResponseWriter, r *http.Request) {

	tokens := r.Header.Get("X-Custom-Token")

	if tokens == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Token received: %s", tokens)
	}

	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		fmt.Fprintln(w, "Content-Type not provided")
	} else {
		fmt.Fprintf(w, "Content-Type: %s\n", contentType)
	}

}

func Forming(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form value", http.StatusBadRequest)
		return
	}

	user := r.Form.Get("username")
	lang := r.Form.Get("language")

	if user == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if lang == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Printf( "Hello %s, you are coding in %s", user, lang)
}

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
	token := r.Header.Get("X-Custom-Token")

	if token == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		fmt.Fprintln(w, "Content-Type not provided")
	} else {
		fmt.Fprintf(w, "Content-Type: %s\n", contentType)
	}

}

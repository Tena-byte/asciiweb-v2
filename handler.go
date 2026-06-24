package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
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

	if user == "" || lang == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")

	if contentType != "application/x-www-form-urlencoded" {
		http.Error(w, "Unsupported Media", http.StatusUnsupportedMediaType)
		return
	}

	fmt.Fprintf(w, "Hello %s, you are coding in %s", user, lang)
}

func StatusCheck(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	if code == "" {
		http.Error(w, "code parameter is required", 400)
		return
	}

	codevalue, err := strconv.Atoi(code)
	if err != nil {
		http.Error(w, "code must be a valid integer", 400)
		return
	}

	if codevalue < 100 || codevalue > 599 {
		http.Error(w, "code must be a valid HTTP status code (100–599)", 400)
		return
	}

	w.WriteHeader(codevalue)

	if codevalue != http.StatusNoContent {
		fmt.Fprintf(w, "Responding with status %d", codevalue)
		return
	}

}

//6

func pingHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "pong")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintln(w, "Greetings, Stranger!")
	} else {
		fmt.Fprintf(w, "Greetings, %s!", name)
	}
}

//7

func render(w http.ResponseWriter, r *http.Request) {

	type PageData struct {
		Title string
		Body  string
	}

	const tmplStr = `
		<!DOCTYPE html>
		<html>
			<head><title>{{.Title}}</title></head>
			<body>
				<h1>{{.Title}}</h1>
				<p>{{.Body}}</p>
			</body>
		</html>
`

	data := PageData{}

	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")

	data.Title = title
	data.Body = body
	
	if title == "" || body == ""{
		http.Error(w, "title and body are required", http.StatusBadRequest)
		return
	}
 
	tmpl := template.Must(template.New("page").Parse(tmplStr))

	w.Header().Set("Content-Type", "text/plain")

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w,  "template execution failed", 500)
	}
}

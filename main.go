package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	mainMux.HandleFunc("/", MethodChecks)
	mainMux.HandleFunc("/echo", echoText)
	mainMux.HandleFunc("/headers", HeadersDetective)
	mainMux.HandleFunc("/form", Forming)
	mainMux.HandleFunc("/status", StatusCheck)
	mainMux.HandleFunc("/render", render)

	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	fmt.Println("server is running on http://localhost:3000")

	if err := http.ListenAndServe(":3000", mainMux); err != nil {
		log.Fatal("server not running")
	}
}

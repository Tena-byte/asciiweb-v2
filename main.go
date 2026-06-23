package main

import (
	"fmt"
	"log"
	"net/http"
)



func main(){


	mux := http.NewServeMux()

	mux.HandleFunc("/", MethodChecks)
	mux.HandleFunc("/echo", echoText)


	fmt.Println("server is running on http://localhost:3000")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal("server not running")
	}
}
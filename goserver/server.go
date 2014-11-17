package main

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":5002", nil)
}

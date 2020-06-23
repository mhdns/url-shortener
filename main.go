package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":5000", nil)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

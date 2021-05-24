package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "hello\n")
}

func callback(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/callback", callback)

	http.ListenAndServe(":8000", nil)
}

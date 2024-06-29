package main

import (
	"fmt"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("go works", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	http.HandleFunc("/go/work", Hello)
	http.ListenAndServe(":8081", nil)
}
